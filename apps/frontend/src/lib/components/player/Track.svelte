<script lang="ts">
	import { createSlider, melt } from '@melt-ui/svelte';

	const {
		elements: { root, range, thumbs },
	} = createSlider({
		defaultValue: [0],
		min: 0,
		max: 100,
		step: 1,
	});

	export let value: number = 0;
</script>

<span use:melt="{$root}" class="relative flex h-[20px] w-full items-center">
	<span class="h-[3px] w-full bg-black/40">
		<span use:melt="{$range}" class="h-[3px] bg-white"></span>
	</span>

	<span
		use:melt="{$thumbs[0]}"
		class="h-1 w-1 active:h-3 hover:h-3 active:w-3 hover:w-3 rounded-full bg-white focus:ring-4 focus:!ring-black/40"
	></span>
</span>
