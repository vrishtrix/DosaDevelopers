<script lang="ts">
	export let contactName: string;
	export let lastText: string;
	export let unreadTexts: number;
</script>

<div class="flex items-center gap-4 border-b py-4 border-brightGrey">
	<div
		class="flex h-12 w-12 items-center justify-center rounded-full bg-neutral-100"
	>
		<img alt="Avatar" class="h-full w-full rounded-[inherit]" />
	</div>

	<div>
		<p class="text-white text-xl">{contactName}</p>
		<p class="text-darkWhite">{lastText}</p>
	</div>

	{#if unreadTexts != 0}
		<div
			class="relative ml-auto rounded-full bg-partyPurple aspect-square h-1/2 text-spaceGrey mr-4"
		>
			<span
				class="absolute top-1/2 left-1/2 translate-x-[-50%] translate-y-[-50%]"
			>
				{unreadTexts}</span
			>
		</div>
	{/if}
</div>
