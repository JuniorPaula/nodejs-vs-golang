import { faker } from "@faker-js/faker"
import { createWriteStream } from "node:fs"

const file1 = createWriteStream("../database/file1.ndjson")
const file2 = createWriteStream("../database/file2.ndjson")
const file3 = createWriteStream("../database/file3.ndjson")

const generateUser = () => {
    return {
        userId: faker.datatype.uuid(),
        username: faker.internet.userName(),
        email: faker.internet.email(),
        phone: faker.phone.number(),
        registerAt: faker.date.past(),
    }
}
;
[file1, file2, file3].forEach((file, index) => {
    const currentFile = `file${index + 1}`
    console.time(currentFile)
    for (let i = 0; i < 1e4; i++) {
        const user = generateUser()
        file.write(JSON.stringify(user) + "\n")
    }
    file.end()
    console.timeEnd(currentFile)
})