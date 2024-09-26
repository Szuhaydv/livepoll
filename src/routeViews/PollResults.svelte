<script>
	import Note from "../components/Note.svelte";
	const options = [];
	const percentages = [];
	$: duration = 5;
	let intervalRef = null;
	function tick() {
		intervalRef = setInterval(() => {
			if (duration == 0 && intervalRef) {
				clearInterval(intervalRef);
				return;
			}
			duration -= 1;
		}, 1000);
	}
	tick();
</script>

<h1 class="text-8xl text-yellow-400 text-center pt-[4vh] font-actionJackson">
	Livepoll
</h1>
<Note title="What exercise should I do?" titleMargin={2}>
	<div class="w-full flex justify-center items-center gap-5 mb-8">
		<img src="/public/assets/clock.svg" alt="Clock icon" />
		<p>
			{Math.floor(duration / 60)}:{String(duration % 60).padStart(2, "0")}
		</p>
	</div>
	<ul class="flex flex-col">
		{#each options as option, index}
			<li>
				<p
					class="flex items-center border-t-2 h-16 border-slate-400 w-full pl-12"
				>
					{option}
				</p>
				<div
					class="flex items-center border-t-2 h-16 border-slate-400 w-full pl-12 {index ==
					option.length - 1
						? 'border-b-2'
						: ''}"
				>
					<div
						class="loading-bar rounded-full h-8 mr-4 border border-black"
						style="width: {percentages[index]}%"
					></div>
					<span class="text-nowrap">{percentages[index]} %</span>
				</div>
			</li>
		{/each}
	</ul>
</Note>

<style>
	.loading-bar {
		background-image: linear-gradient(
			-45deg,
			#ffc33c 0%,
			#ffc33c 25%,
			#ffaf2d 25%,
			#ffaf2d 50%,
			#ffc33c 50%,
			#ffc33c 75%,
			#ffaf2d 75%,
			#ffaf2d 100%
		);
		background-size: 2rem;
		animation: progress 0.75s linear infinite;
		box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
	}

	@keyframes progress {
		to {
			background-position-x: 2rem;
		}
	}
</style>
