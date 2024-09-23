<script>
	export let url;
	export let text;
</script>

<a
	href={url}
	class="text-white visited:text-white hover:no-underline rounded-3xl border-4 border-white bg-gradient-to-b from-lime-300 to-green-700 w-max font-caveatBrush px-16 py-3 text-[2.5rem] shadow-lg focus:border-white active:scale-105 transition"
>
	{text}
</a>
