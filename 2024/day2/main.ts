import input from "./input"

function log(...args: any[]) {
    console.log(...args)
    LOG.write(args.join(" ") + "\n")
}

function group(...args: any[]) {
    console.group(...args)
    LOG.write(args.join(" ") + "\n")
}

function groupEnd() {
    console.groupEnd()
    LOG.write("\n")
}
    

function assertTrue(condition: boolean, ...message: string[]): asserts condition {
    if (!condition) throw new Error(message.join(" "))
}


function loadArrays(rows: string[]) {
    let arr: number[][] = []
    rows.forEach(row => {
        if (row == "") return
        const elems = row.split(" ").map(Number)
        arr.push(elems)
    })

    return arr
}

type Row = {
    safe: boolean,
    data: number[],
    errors: number,
    dir: 'inc' | 'dec'
}

function getRowDirection(row: number[]): "inc" | "dec" {
    let [ups,downs] = [0,0]
    for (let i = 0; i < row.length -1; i++) {
        if (row[i+1] > row[i]) {
            ups++
        } else if (row[i+1] < row[i]) {
            downs++
        }
    }
    return ups > downs ? "inc" : "dec"
}

function getGeneralSafety(row: number[], dir: "inc" | "dec"): number {
    let errCount = 0
    for (let i = 0; i < row.length -1; i++) {
        assertTrue(row[i] != undefined && row[i+1] != undefined, `Invalid elements: ${row[i]} ${row[i+1]}`)
        const diff = row[i+1] - row[i]
        if (Math.abs(diff) > 3) {
            errCount += 1
        } else if (diff == 0) { 
            errCount += 1
        } else if (dir == "inc" && diff < 0) {
            errCount += 1
        } else if (dir == "dec" && diff > 0) {
            errCount += 1
        }
    }
    
    return errCount
}

function parseRow(idx: number, row: Row, canFail: boolean): Row {
    assertTrue(row.data.length >= 3, `Row must have at least 3 elements: [${row.data}]` )

    const { data } = row;

    for (let i = 0; i < data.length - 1; i++) {
        let [e1,e2, ...rest] = data.slice(i) 
        assertTrue(e1 != undefined && e2 != undefined, `Invalid elements: ${e1} ${e2}`)

        if (!validateSafety(e1,e2, row.dir)) {
            if (!canFail) {
                log(`FAILED: ${e1} -> ${e2}`)
                row.safe = false
                return row
            }

            // Remove the problematic element

            if (rest.length == 0) {
                log(`Removing ${e2} from [${data}]`)
                row.data = [...data.slice(0,i+1)]
                return parseRow(idx, row, false)
            }
            
            const e3 = rest[0]
            group(`Checking ${e1} -> ${e3}`)
            if (validateSafety(e1,e3,row.dir)) {
                log(`Removing ${e2} from [${data}]`)
                row.data = [...data.slice(0,i),e1,...rest]
                const res = parseRow(idx, row, false)
                groupEnd()
                return res;
            }
            groupEnd()


            if (validateSafety(e2,e3,row.dir)) {
                log(`Removing ${e1} from [${data}]`)
                row.data = [...data.slice(0,i),e2,...rest]
                return parseRow(idx, row, false)
            }

            log(`FAILED: ${e1} -> ${e2} -> ${e3}`)
            row.safe = false
            return row
        }
    }
    // if (errCount > 1) {
    //     log(`=============== SUCCESSFUL YET UNSAFE: ${errCount} errors, dir: ${dir} =======================`)
    // }
    return row;
}

function validateSafety(r1: number, r2: number, direction: "inc" | "dec") {
    const diff = r2 - r1
    if (direction == "inc" && diff <= 0) {
        log(`Should be increasing: ${r1} -> ${r2}`)
        return false 
    } else if (direction == "dec" && diff >= 0) {
        log(`Should be decreasing: ${r1} -> ${r2}`)
        return false
    }
    
    if (Math.abs(diff) > 3) {
        log(`Should be diff < 3: |${r1} - ${r2}| == ${diff}`)
        return false
    }
    return true
}

function checkRow(r: number[]): { safe: boolean, dir: "inc" | "dec" } {
    const diffs = r.map((e,i) => r[i+1] - e).filter(e => !isNaN(e))

    let dir = diffs.reduce((sum, val) => sum + val, 0) > 0 ? "inc" : "dec" as "inc" | "dec"

    if (!diffs.every((v) => v > 0) && !diffs.every((v) => v < 0)) {
        return {safe: false, dir };
    }

    if (diffs.some(d => Math.abs(d) > 3)) {
        return {safe: false, dir };

    }

    return {safe: true, dir };

}


function solvePart1(rows: number[][]) {
    let safeRows = rows.filter(r => checkRow(r).safe).length

    log("Part 1: " + safeRows + " safe rows")
}



function solvePart2(nums: number[][]) {
    const safeRows = nums.filter(row => {
        const res = checkRow(row);
        if (!res.safe) {
            for (let i = 0; i < row.length; i++) {
                if (checkRow(row.toSpliced(i,1)).safe) {
                    // Removing 1 element makes the row safe
                    return true
                }
            }
        }
        return res.safe
    }).length

    log("Part 2: " + safeRows + " safe rows")

    return safeRows
}


/**
 * 
 * 
 * Main 
 * 
 * 
 */
async function main() {
    // const testInput = "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n"
    // const rows = loadArrays(testInput.split("\n"))
    const rows = loadArrays(input.split("\n"))
    solvePart1(rows)
    solvePart2(rows)
}

const LOG = Bun.file("output.log", {
    type: "text",
}).writer();

main().then(() => {
    LOG.end();
    process.exit(0)
})

// 299 - PART 1

// PART 2
// 353 - WRONG
// 359 too low

// 485 - WRONG
// 511 - WRONG
// 637 - WRONG 

// 712 too high