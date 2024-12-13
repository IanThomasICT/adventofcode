console.log("Hello via Bun!");
const sampleInput = `125 17`;

function parseInput(input?: string) {
    if (!input) {
        input = sampleInput;
    }

    return input.split(' ').map(Number);
}

function solve(input?: string) {
    const performanceStart = performance.now();

    const stones = parseInput(input);

    const dp = new Map<number, Map<number, number>>(); // <stone value, num stones at each depth>

    const blink = (stone: number) => {
        const stones = [stone];
        const valueAsString = `${stone}`;

        if (stone === 0) {
            stones[0] = 1;
        } else if (valueAsString.length % 2 === 0) {
            stones[0] = +valueAsString.slice(0, valueAsString.length / 2);
            stones.push(+valueAsString.slice(valueAsString.length / 2));
        } else {
            stones[0] *= 2024;
        }

        return stones;
    };

    const processStone = (stone: number, absoluteDepth: number, relativeDepth: number, targetDepth: number) => {
        const dpMap = dp.get(stone) ?? new Map<number, number>();

        // if we've already calculated this stone, for this many levels deep, just return that calculation
        if (dpMap.has(targetDepth)) {
            return dpMap.get(targetDepth)!;
        }

        let numNewStones = 0;

        const newStones = blink(stone);
        // blinking generated a new stone
        if (newStones.length > 1) {
            numNewStones++;
        }

        if (relativeDepth < targetDepth - 1) {
            numNewStones += processStone(newStones[0], absoluteDepth + 1, relativeDepth, targetDepth - 1);
            if (newStones.length > 1) {
                numNewStones += processStone(newStones[1], absoluteDepth + 1, relativeDepth, targetDepth - 1);
            }
        }

        // remember how many new stones we found for this stone at this many levels deep
        dpMap.set(targetDepth, numNewStones);
        dp.set(stone, dpMap);
        // console.log('stone', stone, 'at depth', absoluteDepth, 'has', numNewStones, 'new stones beneath it');

        return numNewStones;
    }

    const targetDepth = 75;
    let numStones = stones.length;
    for (const stone of stones) {
        numStones += processStone(stone, 0, 0, targetDepth);
    }

    return {
        performance: performance.now() - performanceStart,
        result: numStones
    }
}


console.log(solve("64599 31 674832 2659361 1 0 8867 321"))
