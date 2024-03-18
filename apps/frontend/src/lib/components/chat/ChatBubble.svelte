<script lang="ts">
	export let type: 'sent' | 'received' | undefined;
	export let message: string | undefined;
</script>

<div
	class="relative flex items-end {type === 'sent'
		? 'ml-auto'
		: null} max-w-prose p-4 {type === 'sent'
		? 'bg-brightGrey'
		: 'bg-partyPurple'} text-white rounded mb-4 {type === 'sent'
		? 'bubble-right'
		: 'bubble-left'}"
>
	<span>
		{message}
	</span>
	<span class="text-xs text-darkWhite ml-auto">11.00</span>
</div>

<style>
	.bubble-right::after {
		content: '';
		position: absolute;
		top: 70%;
		left: 100%;
		height: 10px;
		width: 10px;
		background-color: #ebebf54d;
		clip-path: polygon(0 0, 0% 100%, 100% 50%);
	}

	.bubble-left::before {
		content: '';
		position: absolute;
		top: 70%;
		left: -10px;
		height: 10px;
		width: 10px;
		background-color: #785beb;
		clip-path: polygon(100% 0, 100% 100%, 0 50%);
	}
</style>
