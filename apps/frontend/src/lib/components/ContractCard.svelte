<script lang="ts">
	import GreenCheckIcon from '../../assets/icons/green-check.svg';
	import RedCrossIcon from '../../assets/icons/red-cross.svg';

	export let contactName: string;
	export let duration: string;
	export let date: string;
	export let status: 'signed' | 'cancelled';

	const iconMap: Record<string, string> = {
		signed: GreenCheckIcon,
		cancelled: RedCrossIcon,
	};
</script>

<div class="flex items-center gap-4 border-b py-4 border-brightGrey">
	<div
		class="flex h-12 w-12 items-center justify-center rounded-full bg-neutral-100"
	>
		<img alt="Avatar" class="h-full w-full rounded-[inherit]" />
	</div>

	<div>
		<p class="text-white text-xl">{contactName}</p>
		<p class="text-darkWhite">{duration} • {date}</p>
	</div>

	<div
		class="relative ml-auto rounded-full aspect-square h-1/2 text-spaceGrey mr-4"
	>
		<img src="{iconMap[status]}" alt="Contract status icon" />
	</div>
</div>
