<script>
	export let title;
	export let titleMargin = 10;
	export let isTitleEditable = false;
	let titleRef;
</script>

<div class="flex flex-col gap-9 items-center justify-center">
	<main class="bg-yellow-paper w-[48rem] shadow-2xl py-12 px-24">
		{#if isTitleEditable}
			<div
				class="flex w-full items-center"
				style="margin-bottom: {titleMargin / 4}rem"
			>
				<input
					value={title}
					type="text"
					class="text-ellipsis overflow-hidden whitespace-nowrap bg-transparent w-full border-none focus:outline-none font-bangers text-5xl text-center mr-4"
					bind:this={titleRef}
				/>
				<button
					class="border-none h-10 w-10"
					on:click={() => titleRef.focus()}
				>
					<img
						src="/public/assets/edit.svg"
						alt="Paper and pencil icon"
						class="w-10 h-10"
					/>
				</button>
			</div>
		{:else}
			<h2
				class="font-bangers text-5xl text-center"
				style="margin-bottom: {titleMargin / 4}rem"
			>
				{title}
			</h2>
		{/if}
		<slot></slot>
	</main>
	<slot name="just-below"></slot>
</div>

<style>
	div {
		flex: 1;
	}
</style>
