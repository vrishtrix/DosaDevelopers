<script lang="ts">
	import {
		status,
		isPlaying,
		audioPlayer,
		index,
		trackList,
		addTrack,
	} from '../../store/player';
	import PlayButtonIcon from '../../../assets/icons/play-button.svg';

	export let controls = false;
	export let track = false;
	export let title = '';
	export let artist = '';
	export let file = '';

	const playTrack = () => {
		$audioPlayer.play();
		$isPlaying = true;
	};

	const pauseTrack = () => {
		$audioPlayer.pause();
		$isPlaying = false;
	};

	const loadTrack = ($index: number) => {
		title = $trackList[$index].title;
		artist = $trackList[$index].artist;
		$audioPlayer.src = $trackList[$index].file;
		$audioPlayer.load();
	};

	const addAndPlayTrack = () => {
		addTrack({ title, artist, file });
		$index = $trackList.length - 1;
		loadTrack($index);
		playTrack();
	};
</script>

{#if controls}
	{#if $isPlaying === false}
		<button class="play-button controls" on:click="{playTrack}">
			<img src="{PlayButtonIcon}" alt="play button" />
		</button>
	{:else if $isPlaying === true && ($status === 'waiting' || $status === 'loading' || $status === 'can play some' || $status === 'can play all')}
		<button class="play-button controls" on:click="{pauseTrack}">
			<img src="{PlayButtonIcon}" alt="play button" />
		</button>
	{:else if $isPlaying === true}
		<button class="play-button controls" on:click="{pauseTrack}">
			<img src="{PlayButtonIcon}" alt="play button" />
		</button>
	{/if}
{:else if track}
	{#if title !== $trackList[$index].title}
		<button class="play-button track" on:click="{addAndPlayTrack}">
			<img src="{PlayButtonIcon}" alt="play button" />
		</button>
	{:else if title === $trackList[$index].title && $isPlaying === true && ($status === 'loading' || $status === 'can play some' || $status === 'can play all' || $status === 'waiting')}
		<button class="play-button track playing" on:click="{pauseTrack}">
			<img src="{PlayButtonIcon}" alt="play button" />
		</button>
	{:else if title === $trackList[$index].title && $isPlaying === true}
		<button class="play-button track playing" on:click="{pauseTrack}">
			<img src="{PlayButtonIcon}" alt="play button" />
		</button>
	{:else if title === $trackList[$index].title && $isPlaying === false}
		<button class="play-button track playing" on:click="{playTrack}">
			<img src="{PlayButtonIcon}" alt="play button" />
		</button>
	{/if}
{/if}
