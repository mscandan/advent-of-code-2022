const path = require('path');
const fs = require('fs');

const input = fs
	.readFileSync(path.join(__dirname, 'problem.input'), 'utf8')
	.toString()
	.trim()
	.split('\n')
	.map((line) => line.split(''));

let S;
let E;

for (let y = 0; y < input.length; y++) {
	for (let x = 0; x < input[y].length; x++) {
		const cell = input[y][x];
		if (cell === 'S') {
			S = { x, y };
			input[y][x] = 'a';
		} else if (cell === 'E') {
			E = { x, y };
			input[y][x] = 'z';
		}
		input[y][x] = input[y][x].charCodeAt(0);
	}
}

const toId = (x, y) => `${x},${y}`;

function getNeighbors(x, y) {
	return [
		{ x: x, y: y - 1 },
		{ x: x - 1, y: y },
		{ x: x + 1, y: y },
		{ x: x, y: y + 1 },
	].filter((coord) => typeof input[coord.y]?.[coord.x] !== 'undefined');
}

function buildFrontier(from_x, from_y) {
	const frontier = [];
	frontier.push({ x: from_x, y: from_y });

	const came_from = new Map();
	came_from.set(toId(from_x, from_y), null);
	while (frontier.length > 0) {
		const current = frontier.shift();
		const current_val = input[current.y][current.x];

		let neighbors = getNeighbors(current.x, current.y);
		for (let next of neighbors) {
			const next_cell = input[next.y][next.x];
			const next_id = toId(next.x, next.y);

			if (next_cell - current_val > 1 || came_from.has(next_id)) {
				continue;
			}

			// Coord is walkable
			const current_id = toId(current.x, current.y);
			frontier.push(next);
			came_from.set(next_id, current_id);
		}
	}

	return came_from;
}

function getShortestPath(from_x, from_y, to_x, to_y) {
	const from_id = toId(from_x, from_y);
	const to_id = toId(to_x, to_y);
	const came_from = buildFrontier(from_x, from_y);
	let current = to_id;

	let path = [];
	while (current !== undefined && current !== from_id) {
		path.push(current);
		current = came_from.get(current);
	}

	if (current === undefined) {
		return [];
	}

	path.reverse();

	return path;
}

const paths = getShortestPath(S.x, S.y, E.x, E.y);
console.log("Solution for Day 12 Part 1: ", paths.length)