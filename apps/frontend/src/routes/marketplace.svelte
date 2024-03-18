<script lang="ts">
	import AppShell from './../lib/components/AppShell.svelte';
	import CategoryBackground from '../assets/category-bg.png';

	const categories = ['Events Near You', 'Local Gigs', 'Competitions'];
	const events = [
		{
			title: 'Bands & Groups',
			description: 'Blues bands, Mariachis',
		},
		{
			title: 'Ensembles',
			description: 'Chamber Orchestras',
		},
		{
			title: 'Singers',
			description: 'Country Singers, Rappers',
		},
		{
			title: 'Guitarists',
			description: 'Electric Guitarists',
		},
		{
			title: 'DJs',
			descriptions: 'Bar Mitzvah DJ',
		},
	];
</script>

<AppShell>
	<div>
		<h1 class="text-2xl mb-2 text-white font-ownersTrialWide">
			Marketplace
		</h1>
		<p class="text-darkWhite">
			Get into the groove and never miss a beat with unmute's guide to the
			hottest gigs, concerts, and jam sessions happening in your
			neighbourhood.
		</p>
	</div>

	<div class="overflow-x-scroll w-full snap-x snap-mandatory flex gap-4 mt-4">
		{#each categories as category}
			<div
				class="p-2 border border-brightGrey bg-[#252525CC] rounded-xl h-[150px] aspect-video"
			>
				<div
					class="flex flex-1 items-end p-4 h-full w-full rounded-lg"
					style="background: url({CategoryBackground}) no-repeat; background-size: 300px auto;"
				>
					<h2 class="text-white text-xl">{category}</h2>
				</div>
			</div>
		{/each}
	</div>

	<div class="mt-10">
		<h1 class="text-2xl mb-2 text-white font-ownersTrialWide">
			Events near you
		</h1>
		<p class="text-darkWhite">
			Curated list events around you to take part in and network with the
			local musicians!
		</p>
	</div>

	<div
		class="overflow-x-scroll w-[96%] snap-x snap-mandatory flex gap-4 mt-4"
	>
		{#each events as event}
			<div
				class="flex flex-col justify-end p-4 rounded-xl h-[150px] aspect-video"
				style="background: linear-gradient(
          			rgba(0, 0, 0, 0.7), 
          			rgba(0, 0, 0, 0.7)
       			), url('https://source.unsplash.com/random') no-repeat; background-size: 300px auto;"
			>
				<h2 class="text-white text-lg">{event.title}</h2>
				<p class="text-darkWhite text-sm">{event.description}</p>
			</div>
		{/each}
	</div>
</AppShell>
