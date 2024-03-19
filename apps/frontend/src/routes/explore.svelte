<script lang="ts">
	import AppShell from '../lib/components/AppShell.svelte';
	import CategoryBackground from '../assets/category-bg.png';

  import WelcomeAditya from '../assets/icons/Welcome.png';
  import Play from '../assets/icons/play.png';
	import Tile1 from '../assets/icons/tile1.png';
	import Tile2 from '../assets/icons/tile2.png';
	import Tile3 from '../assets/icons/tile3.png';
	import Tile4 from '../assets/icons/tile4.png';

	const genres = ['Techno', 'Rap', 'Afro Beats', 'Alternate Rock', 'Indie Rock'];
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
  <!--Genres-->
  <div class="flex w-full items-center">
    <img
      src="{WelcomeAditya}"
      alt="Welcome Aditya"
      class="mr-4 mb-4"
    />
    <h2 class="text-white text-3xl font-semibold mb-4">
      Explore Genres
    </h2>
  </div>
  <p class="text-gray-300 text-xl mb-4 ml-4">
    Your Music Management Hub! Elevate your career with
    analytics, gigs, and fan connections. Let's make music
    magic!
  </p>
<div class="overflow-x-scroll w-full snap-x snap-mandatory flex gap-4 mt-4 no-scrollbar mb-10">
  
  {#each genres as genres}
    <div
      class="p-2 border border-brightGrey bg-[#252525CC] rounded-xl h-[150px] aspect-video"
    >
      <div
        class="flex flex-1 items-end p-4 h-full w-full rounded-lg"
        style="background: url({CategoryBackground}) no-repeat; background-size: 300px auto;"
      >
        <h2 class="text-white text-xl">{genres}</h2>
      </div>
    </div>
  {/each}
</div>

<!--Artists-->
<div class="flex w-full items-center">
  <img
    src="{WelcomeAditya}"
    alt="Welcome Aditya"
    class="mr-4 mb-4"
  />
  <h2 class="text-white text-3xl font-semibold mb-4">
    Explore Artists
  </h2>
</div>
<p class="text-gray-300 text-xl mb-4 ml-4">
  Escape the ordinary. Explore extraordinary music beyond the mainstream. Listen and get them discovered.
</p>
<div class="overflow-x-scroll w-full snap-x snap-mandatory flex gap-4 mt-4 no-scrollbar mb-10">
  
</div>

<!--Curated Media-->

</AppShell>

<style>
	.no-scrollbar::-webkit-scrollbar {
		display: none;
	}
	.no-scrollbar {
		-ms-overflow-style: none;
		scrollbar-width: none;
	}
</style>