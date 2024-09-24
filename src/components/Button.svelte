<script>
	export let url;
	export let text;
	export let actionButton = false;
	export let action;
	const styles =
		"text-white visited:text-white hover:no-underline rounded-3xl border-white border-4 bg-gradient-to-b from-lime-300 to-green-700 w-max px-16 inline-block shadow-lg focus:border-white active:scale-95 transition";
</script>

{#if actionButton}
	<button class={styles} on:click={action}>{text}</button>
{:else}
	<a href={url} class={styles}>
		{text}
	</a>
{/if}
