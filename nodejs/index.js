import { createWriteStream } from "node:fs"
import { PassThrough, Readable } from "node:stream"
import { readdir } from "node:fs/promises"
import { fork } from "node:child_process"
import { pipeline } from "node:stream/promises"

const outputFilename = "../database/output-gmail.ndjson"
const backgroundPath = "./backgroundJob.js"
const output = createWriteStream(outputFilename)

console.time("child_process")

function mergeStreams(streams) {
    let pass = new PassThrough()
    let waiting = streams.length
    for (let stream of streams) {
        pass = stream.pipe(pass, { end: false })
        stream.once("end", () => --waiting === 0 && pass.emit("end"))
    }
    return pass
}

function childProcessAsStream(cp, file) {
    const stream = Readable({
        read() {}
    })

    cp.on("message", ({ status, message}) => {
        if (status === "error") {
            console.error({
                msg: "Got an error!",
                pid: cp.pid,
                message: message.split("\n")
            })
            stream.push(null)
            return
        }
        stream.push(JSON.stringify({
            ...message,
            file,
            pid: cp.pid
        }).concat("\n"))
    })

    cp.send(file)

    return stream
}

const files = (await readdir("../database")).filter(item => !item.includes("output"))

const counters = {}
const childProcesses = []

for (const file of files) {
    const cp = fork(backgroundPath, [], {
        silent: false,
    })
    counters[cp.pid] = { counter: 1 }

    const stream = childProcessAsStream(cp, `../database/${file}`)
    childProcesses.push(stream)
}

const allStreams = mergeStreams(childProcesses)
await pipeline(
    allStreams,
    async function* (source) {
        for await (const chunk of source) {
            for (const line of chunk.toString().trim().split("\n")) {
                const { file, ...data } = JSON.parse(line)
                const counter = counters[data.pid].counter++

                console.log(`${file} found ${counter} so far...`)
                yield JSON.stringify(data).concat("\n")
            }
        }
    },
    output
)

console.timeEnd("child_process")