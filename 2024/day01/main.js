
function loadAndSortArrays(rows) {
    let left = []; let right = [];
    rows.forEach(row => {
        if (row == "") return
        const [l,r] = row.split("   ").map(Number)
        left.push(l)
        right.push(r)
    })
    left.sort((a,b) => a-b)
    right.sort((a,b) => a-b)

    return [left, right]
}

function solvePart1(left, right) {
    let sum = 0
    for (let i = 0; i < left.length; i++) {
        sum += (Math.abs(left[i] - right[i]))
    }

    console.log("Part 1: " + sum)
}

function solvePart2(left, right) {
    let sum = 0
    for (let i = 0; i < left.length; i++) {
        const c = right.filter(e => e == left[i]).length
        sum += (left[i] * c)
    }

    console.log("Part 2: " + sum)
}


function main() {
    const rows = document.querySelector("pre").innerText.trim().split("\n")
    const [left, right] = loadAndSortArrays(rows)
    solvePart1(left, right)
    solvePart2(left, right)
}

main()