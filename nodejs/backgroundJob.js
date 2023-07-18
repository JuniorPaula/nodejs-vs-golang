import { createReadStream } from "node:fs"
import { pipeline } from "node:stream/promises"
import split from "split"

console.log("start background jobs!", process.pid)

process.once("message", async (message) => {
    try {
        await pipeline(
            createReadStream(message),
            split(),
            async function* (source) {
                for await (const chunk of source) {
                    if (!chunk.length) continue

                    const item = JSON.parse(chunk)
                    if (!item.email.includes("gmail")) continue

                    process.send({
                        status: "success",
                        message: item
                    })
                }
            }
        )
    } catch (error) {
        process.send({
            status: "error",
            message: error.message
        })
    }
})
