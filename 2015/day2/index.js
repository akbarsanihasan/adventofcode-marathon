import fs from "node:fs/promises";
import path from "node:path";

const inputFile = await fs.readFile(
    `${path.dirname(process.argv[1])}/input.txt`,
    "utf8",
);
const inputParsed = inputFile
    .trim()
    .split(/\n/g)
    .map((input) => input.split("x").map((n) => parseInt(n)));

let firstProblem = 0;
for (const [length, width, height] of inputParsed) {
    const areax = length * width;
    const areay = width * height;
    const areaz = height * length;
    const smallest = [areax, areay, areaz].sort((a, b) => a - b)[0];
    firstProblem += 2 * areax + 2 * areay + 2 * areaz + smallest;
}

let secondProblem = 0;
for (const [length, width, height] of inputParsed) {
    const [x, y] = [length, width, height].sort((a, b) => a - b);
    secondProblem += x * 2 + y * 2 + length * width * height;
}

console.log(firstProblem);
console.log(secondProblem);
