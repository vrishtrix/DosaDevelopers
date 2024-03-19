<script lang="ts">
	import AlbumCover from '../assets/album-cover.png';
	import IconHeart from '../assets/icons/heart.svg';
	import PlayerCard from '../assets/player-card.svg';
</script>

<div class="h-[40vh] flex justify-between items-end py-4 px-8 bg-green-900">
	<div class="flex items-end gap-4 h-3/4">
		<img
			src="{AlbumCover}"
			alt="Album Cover"
			class="h-full aspect-square"
		/>
		<div>
			<h1 class="text-3xl text-white">Zakir</h1>
			<p class="text-lg text-darkWhite">Nalaayak</p>
		</div>
	</div>

	<div class="flex flex-row items-center">
		<img src="{IconHeart}" alt="Heart Icon" />
		<span class="mx-2 text-2xl text-darkWhite">25k</span>
		<button
			class="px-4 py-2 mx-2 font-bold uppercase rounded bg-partyPurple text-spaceGrey"
		>
			Chat
		</button>
		<button
			class="px-4 py-2 mx-2 font-bold uppercase rounded bg-partyPurple text-spaceGrey"
		>
			Follow
		</button>
	</div>
</div>

<div class="flex p-8">
	<div class="w-2/3">
		<h1 class="mb-2 text-2xl text-white font-ownersTrialWide">Threads</h1>
		<p class="text-darkWhite">(No feedbacks for this track yet.)</p>
	</div>

	<div class="flex flex-col w-1/3">
		<h1 class="mb-4 text-2xl text-white font-ownersTrialWide">
			More From Naayak
		</h1>
		<img src="{PlayerCard}" alt="Music Card" />
		<img src="{PlayerCard}" alt="Music Card" />
		<img src="{PlayerCard}" alt="Music Card" />
		<img src="{PlayerCard}" alt="Music Card" />
	</div>
</div>
