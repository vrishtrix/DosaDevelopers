<script lang="ts">
	import GreenCheckIcon from '../../assets/icons/green-check.svg';
	import RedCrossIcon from '../../assets/icons/red-cross.svg';
	import YellowCheckIcon from '../../assets/icons/yellow-check.svg';

	export let contactName: string;
	export let duration: string;
	export let date: string;
	export let status: 'signed' | 'cancelled';

	const iconMap: Record<string, string> = {
		signed: GreenCheckIcon,
		cancelled: RedCrossIcon,
		progress: YellowCheckIcon,
	};
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
		<p class="text-darkWhite">{duration} • {date}</p>
	</div>

	<div
		class="relative ml-auto mr-4 rounded-full aspect-square h-1/2 text-spaceGrey"
	>
		<img src="{iconMap[status]}" alt="Contract status icon" />
	</div>
</div>
