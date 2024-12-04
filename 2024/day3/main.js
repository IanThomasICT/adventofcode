


function solvePart1(input) {
    valid = [...input.matchAll(/mul\(\d{1,3},\d{1,3}\)/g)]
    let mults = valid.map(m => m[0].slice(4,-1).split(",").map(Number))
    
    let sum = 0
    mults.forEach(m => sum += (m[0]*m[1]))

    console.log("Part 1: " + sum)
}

function solvePart2(input) {
    let validMults = input.split("do()").map(s => s.split("don't()")[0]).flatMap(l => {
        // valid rows
        let ops = [...l.matchAll(/mul\(\d{1,3},\d{1,3}\)/g)]
        let mults = ops.map(m => m[0].slice(4,-1).split(",").map(Number))
        return mults;
    })
    let sum = validMults.reduce((sum, val) => sum + val[0]*val[1], 0)
    console.log("Part 2: " + sum)
}

function main() {
    let input = document.querySelector("pre").innerText.trim()
    solvePart1(input)
    solvePart2(input)
}

main()