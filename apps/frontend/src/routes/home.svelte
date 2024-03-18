<script lang="ts">
	import { navigate } from 'svelte-routing';

	import UnmuteLogo from '../assets/logo.svg';
	import landingBg from '../assets/landing-bg.png';
</script>

<div
	class="flex flex-col justify-between items-start h-screen w-full p-8 bg-contain bg-center-center"
	style="background: url({landingBg}) no-repeat;"
>
	<img src="{UnmuteLogo}" alt="Unmute Logo" class="h-[75px]" />

	<div class="my-auto">
		<h1 class="text-4xl text-white my-8 font-ownersTrialWide max-w-prose">
			At <span class="text-partyPurple">unmute</span> we're flipping the script
			on the music scene
		</h1>
		<p class="text-darkWhite max-w-prose">
			Tired of the same old mainstream tunes? Dive into our vibrant
			underground music universe, where undiscovered gems sparkle and
			hidden talents shine.
		</p>

		<button
			on:click="{() => navigate('/register')}"
			class="bg-partyPurple text-spaceGrey rounded py-2 px-4 mt-20 font-bold uppercase"
		>
			Get Started
		</button>
	</div>
</div>
