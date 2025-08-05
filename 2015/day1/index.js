import fs from "node:fs/promises";
import path from "node:path";

const inputFile = await fs.readFile(
    `${path.dirname(process.argv[1])}/input.txt`,
    "utf8",
);
const inputParsed = inputFile
    .replaceAll(/\s/g, "")
    .split("")
    .map((input) => {
        return input === "(" ? 1 : input === ")" ? -1 : null;
    });

let currentFloor = 0;
let movePosition = 0;
while (currentFloor >= 0) {
    currentFloor += inputParsed[movePosition];
    movePosition++;
}

console.log(inputParsed.reduce((acc, curr) => acc + curr));
console.log(movePosition);
