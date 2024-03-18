<script lang="ts">
	import Controls from './Controls.svelte';
	import Track from './Track.svelte';

	export let title: string;
	export let artists: string;
</script>

<div
	class="grid grid-cols-5 gap-4 items-center border border-brightGrey rounded-full w-full bg-[#42424266] px-8 py-2 my-2 max-h-[75px]"
>
	<div class="col-span-1">
		<h1 class="text-white truncate">{title}</h1>
		<p class="text-darkWhite truncate">{artists}</p>
	</div>

	<Controls />
</div>
