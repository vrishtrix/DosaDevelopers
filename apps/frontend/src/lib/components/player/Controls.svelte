<script lang="ts">
	import PlayButton from './PlayButton.svelte';
	import Slider from './Track.svelte';
	import {
		status,
		isPlaying,
		audioPlayer,
		index,
		trackList,
	} from '../../store/player';
	import { format } from '../../utilities';
	import { onMount } from 'svelte';

	import AudioPreviousIcon from '../../../assets/icons/audio-previous.svg';
	import Track from './Track.svelte';

	let duration = 0;
	let currentTime = 0;
	let formattedTime = format(currentTime);
	let paused = true;
	let volume = 1;

	let slider: Slider;
	let rAF: any = null;

	let title = $trackList[$index].title;
	let artist = $trackList[$index].artist;
	let src = $trackList[$index].file;

	const whilePlaying = () => {
		slider.value = audio.currentTime;
		currentTime = slider.value;
		rAF = requestAnimationFrame(whilePlaying);
	};

	const loadTrack = ($index: number) => {
		title = $trackList[$index].title;
		artist = $trackList[$index].artist;
		$audioPlayer.src = $trackList[$index].file;
		$audioPlayer.load();
	};

	const playTrack = () => {
		$isPlaying = true;
		requestAnimationFrame(whilePlaying);
		$audioPlayer.play();
	};

	const pauseTrack = () => {
		$isPlaying = false;
		cancelAnimationFrame(rAF);
		$audioPlayer.pause();
	};

	const movePosition = () => {
		let time = slider.value;
		if (!audio.paused) {
			cancelAnimationFrame(rAF);
		}
	};

	const updatePosition = () => {
		audio.currentTime = slider.value;
		if (!audio.paused) {
			requestAnimationFrame(whilePlaying);
		}
	};

	const previousTrack = () => {
		$isPlaying = false;
		currentTime = 0;
		if ($index > 0) {
			$index -= 1;
		} else {
			$index = $trackList.length - 1;
		}
		loadTrack($index);
		playTrack();
	};

	const nextTrack = () => {
		$isPlaying = false;
		currentTime = 0;
		if ($index < $trackList.length - 1) {
			$index += 1;
		} else {
			$index = 0;
		}
		loadTrack($index);
		playTrack();
	};

	onMount(() => {
		$audioPlayer.load();
	});
</script>

<div class="col-span-3 flex gap-4">
	<span class="text-sm text-darkWhite">{format(currentTime)}</span>
	<Track />
	<span class="text-sm text-darkWhite">{format(duration)}</span>
</div>

<div class="col-span-1">
	<div class="flex justify-between gap-4 w-full">
		<audio
			bind:this="{$audioPlayer}"
			bind:duration
			bind:currentTime
			bind:paused
			bind:volume
			on:canplay="{() => ($status = 'can play some')}"
			on:canplaythrough="{() => ($status = 'can play all')}"
			on:waiting="{() => ($status = 'waiting')}"
			on:timeupdate="{() => ($status = 'playing')}"
			on:seeking="{() => ($status = 'seeking')}"
			on:ended="{() => {
				$isPlaying = false;
				currentTime = 0;
			}}"
			{src}
		></audio>

		<button on:click="{previousTrack}">
			<img src="{AudioPreviousIcon}" alt="Previous Track" />
		</button>

		<PlayButton controls />

		<button on:click="{nextTrack}">
			<img
				src="{AudioPreviousIcon}"
				alt="Next Track"
				class="rotate-180"
			/>
		</button>
	</div>
</div>
