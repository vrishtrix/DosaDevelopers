<script lang="ts">
	export let contactName: string;
	export let lastText: string;
	export let unreadTexts: number;
</script>

<div class="flex items-center gap-4 py-4 border-b border-brightGrey">
	<div
		class="flex items-center justify-center w-12 h-12 rounded-full bg-neutral-100"
	>
		<img
			src="https://i.pravatar.cc/300?img=2"
			alt="Avatar"
			class="h-full w-full rounded-[inherit]"
		/>
	</div>

	<div>
		<p class="text-xl text-white">{contactName}</p>
		<p class="text-darkWhite">{lastText}</p>
	</div>

	{#if unreadTexts != 0}
		<div
			class="relative ml-auto mr-4 rounded-full bg-partyPurple aspect-square h-1/2 text-spaceGrey"
		>
			<span
				class="absolute top-1/2 left-1/2 translate-x-[-50%] translate-y-[-50%]"
			>
				{unreadTexts}</span
			>
		</div>
	{/if}
</div>
