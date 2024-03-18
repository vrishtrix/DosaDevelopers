<script lang="ts">
	import EditIcon from '../assets/icons/edit.svg';
	import PlusIcon from '../assets/icons/plus.svg';
	import SendIcon from '../assets/icons/send.svg';
	import MicrophoneIcon from '../assets/icons/microphone.svg';
	import PhoneFilledIcon from '../assets/icons/phone-filled.svg';
	import CameraFilledIcon from '../assets/icons/camera-filled.svg';
	import ContactCard from '../lib/components/chat/ContactCard.svelte';
	import ChatBubble from '../lib/components/chat/ChatBubble.svelte';

	let messageBox: HTMLInputElement;
</script>

<div class="grid grid-cols-3 bg-spaceGrey">
	<div class="flex flex-col gap-4 px-4 py-8 border-r border-brightGrey">
		<div class="flex justify-between items-center mb-8">
			<h1 class="text-white text-3xl text-ownersTrialWide">Chats</h1>
			<img src="{EditIcon}" alt="Edit icon" class="h-8" />
		</div>

		{#each Array(5).fill(0) as item}
			<ContactCard
				contactName="Vispute"
				lastText="are you busy right now?"
				unreadTexts="{2}"
			/>
		{/each}
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

		<div class="py-4 px-8 overflow-y-scroll h-full">
			<ChatBubble
				type="sent"
				message="Hey There! I'm Aditya, nice to meet you. Are you here for the gig too?"
			/>
			<ChatBubble
				type="received"
				message="Hey There! I'm Aditya, nice to meet you. Are you here for the gig too?"
			/>
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
