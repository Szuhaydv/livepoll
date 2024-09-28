//import { writable } from "svelte/store";
//
//function createPoll() {
//	const poll = new Poll("", 0, [], "")
//
//	return writable
//}

export class Poll {
	constructor(title, duration, options, createdAt) {
		this.title = title;
		this.duration = duration;
		this.options = options ? options : [];
		this.createdAt = createdAt;
	}
}

export class Temp {
	constructor() {
		this.temperature = 23
	}
}

export async function fetchPoll(pollID, poll, timer, totalVotes) {
	try {
		const getPollResponse = await fetch("/polls/" + pollID, {
			method: "GET",
		});
		if (!getPollResponse.ok) {
			const errMessage = await getPollResponse.text();
			console.error("Error getting poll: ", errMessage);
			return;
		}

		const data = await getPollResponse.json();
		console.log(data);
		poll.update((poll) => {
			poll.title = data.title;
			poll.duration = data.duration;
			poll.options = data.options;
			poll.createdAt = data.created_at;
			return poll;
		});
		for (const option of data.options) {
			totalVotes.update((n) => n + option.votes);
		}
		calculatePercentages(poll, totalVotes);
		initTimer(timer, poll);
	} catch (error) {
		console.error("Something went wrong: ", error);
	}
}
function initTimer(timer, poll) {
	calculateEndDate(timer, poll);
	timer.update((timerValue) => {
		timerValue.countdownRef = setInterval(() => {
			updateTimeRemaining(timer);
		}, 1000);
		return timerValue;
	});
}
function calculatePercentages(poll, votesRef) {
	let totalVotes
	const unsub = votesRef.subscribe((value) => totalVotes = value)
	unsub()
	const isZero = totalVotes === 0
	poll.update((poll) => {
		for (const option of poll.options) {
			if (isZero) {
				option.percentage = 0;
			} else {
				option.percentage = Math.round(
					(option.votes / totalVotes) * 100,
				);
			}
		}
		return poll;
	});
}
function calculateEndDate(timerRef, pollRef) {
	let createdAt;
	let durationInParts;
	const unsub = pollRef.subscribe((value) => {
		createdAt = new Date(value.createdAt);
		durationInParts = value.duration.split(":").map(Number);
	});
	const durationInMs =
		durationInParts[1] * 60000 + durationInParts[2] * 1000;
	console.log(createdAt, durationInParts, durationInMs);
	timerRef.update((value) => {
		value.inactivationDate = new Date(
			createdAt.getTime() + durationInMs,
		);
		return value;
	});
	unsub();
}
function updateTimeRemaining(ref) {
	ref.update((value) => {
		let duration = Math.round(
			(value.inactivationDate.getTime() - Date.now()) / 1000,
		);
		if (duration < 0) {
			value.timeRemaining = 0;
			if (value.countdownRef) {
				clearInterval(value.countdownRef);
			}
		} else {
			value.timeRemaining = duration;
		}
		return value;
	});
}
