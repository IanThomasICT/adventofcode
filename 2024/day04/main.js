
let input = "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

function solvePart1(rows) {
    let mat = rows.map(r => r.split(""))
    let hOccs = rows.reduce((sum, r) => sum + ([...r.matchAll(/(XMAS)/g)].length + [...r.matchAll(/(SAMX)/g)].length),0)

    let cols = []
    for (let c = 0; c < mat[0].length; c++) {
        cols[c] = mat.map((r,i) => mat[i][c]).join("")
    }
    let vOccs = cols.reduce((sum, c) => sum + ([...c.matchAll(/(XMAS)/g)].length + [...c.matchAll(/(SAMX)/g)].length),0)
    
    let diagonals = getDiagonals(mat)
    let dOccs = diagonals.reduce((sum, d) => sum + ([...d.matchAll(/(XMAS)/g)].length + [...d.matchAll(/(SAMX)/g)].length),0)

    console.log(`XMAS appears: ${hOccs + vOccs + dOccs} times`)
}

function getDiagonals(mat) {
    let diagonals = []
    
    // top left to bottom right
    for (let row = mat.length-1; row >= 0; row--) {
        let i = row
        let j = 0
        let diag = ""
        while (i < mat.length) {
            diag += mat[i][j]
            i++
            j++
        }
        diagonals.push(diag)
    }

    for (let c = 1; c < mat[0].length; c++) {
        let i = c
        let j = 0
        let diag = ""
        while (i < mat[0].length) {
            diag += mat[j][i]
            i++
            j++
        }
        diagonals.push(diag)
    }

    // top right to bottom left
    for (let row = mat.length-1; row >= 0; row--) {
        let i = row
        let j = mat[0].length-1
        let diag = ""
        while (i < mat.length) {
            diag += mat[i][j]
            i++
            j--
        }
        diagonals.push(diag)
    }

    for (let c = mat[0].length-2; c >= 0; c--) {
        let i = 0
        let j = c
        let diag = ""
        while (j >= 0) {
            diag += mat[i][j]
            i++
            j--
        }
        diagonals.push(diag)
    }

    return diagonals;
}

function containsXMAS(mat, i, j) {
    // On an edge, return false
    if (i == 0 || j == 0 || i == mat.length-1 || j == mat[0].length-1 || mat[i][j] != "A") {
        return false
    }

    // MAS || SAM on tl-br diagonal
    if (mat[i-1][j-1] == "M" && mat[i+1][j+1] == "S" || mat[i-1][j-1] == "S" && mat[i+1][j+1] == "M") {
        // MAS || SAM on tr-bl diagonal
        if (mat[i-1][j+1] == "M" && mat[i+1][j-1] == "S" || mat[i-1][j+1] == "S" && mat[i+1][j-1] == "M") {
            return true
        }

        return false
    } 


    return false
}

function solvePart2(rows) {
    let mat = rows.map(r => r.split(""))

    let xOcc = 0
    for (let i = 0; i < mat.length; i++) {
        for (let j = 0; j < mat[0].length; j++) {
            if (mat[i][j] == "A") {
                containsXMAS(mat, i, j) ? xOcc++ : null
            }
        }
    }

    console.log(`X-MAS appears: ${xOcc} times`)
}

async function main() {
    let rows = document.querySelector("pre").innerText.trim().split("\n")
    solvePart1(rows)
    solvePart2(rows)
}

main().then(() => {
    console.log("Finished running main.")
    process.exit(1)
})