<script lang="ts">
	import EditIcon from '../assets/icons/edit.svg';
	import PlusIcon from '../assets/icons/plus.svg';
	import SendIcon from '../assets/icons/send.svg';
	import MicrophoneIcon from '../assets/icons/microphone.svg';
	import PhoneFilledIcon from '../assets/icons/phone-filled.svg';
	import CameraFilledIcon from '../assets/icons/camera-filled.svg';
	import ContactCard from '../lib/components/chat/ContactCard.svelte';
	import ChatBubble from '../lib/components/chat/ChatBubble.svelte';

	import { messages, username } from '../lib/store/user';

	let messageBox: HTMLInputElement;

	const socket = new WebSocket(import.meta.env.VITE_WEBSOCKET_CHAT_URL);

	socket.onopen = () => {
		socket.send(
			JSON.stringify({
				action: 'initial_connection',
				username: $username,
			})
		);
		console.log('Websocket opened successfully!');
	};

	socket.onmessage = (e) => {
		const data = JSON.parse(e.data);

		//@ts-ignore
		messages.update((msgs) => {
			if (Array.isArray(data)) return [...msgs, ...data];
			return [...msgs, data];
		});
	};

	const sendMessage = () => {
		if (messageBox.value.length) {
			socket.send(
				JSON.stringify({
					action: 'post_message',
					message: messageBox.value,
				})
			);

			//@ts-ignore
			messages.update((msgs) => {
				return [
					...msgs,
					{
						username: $username,
						type: 'sent',
						message: messageBox.value,
					},
				];
			});
		}
		messageBox.value = '';
	};
</script>

<div class="grid grid-cols-3 bg-spaceGrey">
	<div class="flex flex-col gap-4 px-4 py-8 border-r border-brightGrey">
		<div class="flex justify-between items-center mb-8">
			<h1 class="text-white text-3xl text-ownersTrialWide">Chats</h1>
			<img src="{EditIcon}" alt="Edit icon" class="h-8" />
		</div>

		<div class="overflow-y-scroll">
			{#each Array(5).fill(0) as item}
				<ContactCard
					contactName="Vispute"
					lastText="are you busy right now?"
					unreadTexts="{2}"
				/>
			{/each}
		</div>
	</div>

	<div class="col-span-2 h-full">
		<div
			class="w-full items-center gap-4 flex p-4 text-ownersTrialWide border-b border-brightGrey"
		>
			<div
				class="flex h-12 w-12 items-center justify-center rounded-full bg-neutral-100"
			>
				<img alt="Avatar" class="h-full w-full rounded-[inherit]" />
			</div>
			<p class="text-white text-xl">Vispute</p>

			<div class="ml-auto flex items-center gap-8">
				<img src="{PhoneFilledIcon}" alt="Phone Icon" />
				<img src="{CameraFilledIcon}" alt="Camera Icon" />
			</div>
		</div>

		<div class="py-4 px-8 overflow-y-scroll messages-height">
			{#each $messages as message}
				<ChatBubble type="{message.type}" message="{message.message}" />
			{/each}
		</div>

		<div
			class="flex mt-auto space-between p-4 border-t border-brightGrey gap-4"
		>
			<button
				class="relative rounded-full bg-partyPurple aspect-square w-10"
			>
				<img
					src="{PlusIcon}"
					alt="Plus Icon"
					class="absolute top-1/2 left-1/2 translate-x-[-50%] translate-y-[-50%]"
				/>
			</button>

			<input
				bind:this="{messageBox}"
				type="text"
				placeholder="Type something..."
				class="border border-brightGrey active:border-partyPurple active:ring-partyPurple bg-transparent py-2 px-4 text-white rounded-full w-full"
			/>

			<img src="{MicrophoneIcon}" alt="Microphone Icon" />

			<button
				on:click="{sendMessage}"
				type="submit"
				class="relative rounded-full bg-partyPurple aspect-square w-10"
			>
				<img
					src="{SendIcon}"
					alt="Send Icon"
					class="absolute top-1/2 left-1/2 translate-x-[-50%] translate-y-[-50%]"
				/>
			</button>
		</div>
	</div>
</div>

<style>
	.messages-height {
		height: calc(100vh - 200px);
	}
</style>
