<script lang="ts">
	import { Link } from 'svelte-routing';
	import { createCollapsible, melt } from '@melt-ui/svelte';
	import { slide } from 'svelte/transition';
	import { signOut } from '../lib/middleware/auth';

	import UnmuteLogo from '../assets/logo.svg';
	import DashboardIcon from '../assets/icons/dashboard.png';
	import DraftsIcon from '../assets/icons/drafts.png';
	import ChatsIcon from '../assets/icons/chats.png';
	import ExploreIcon from '../assets/icons/explore.png';
	import WebsiteBuilderIcon from '../assets/icons/websiteBuilder.png';
	import PaymentsIcon from '../assets/icons/payments.png';
	import ContractsIcon from '../assets/icons/contracts.png';
	import ProfileIcon from '../assets/icons/profile.png';
	import SettingsIcon from '../assets/icons/settings.png';
	import RevenueGraph from '../assets/icons/RevenueGraph.png';
	import AudienceGraph from '../assets/icons/AudienceGraph.png';
	import FanGraph from '../assets/icons/FanGraph.png';
	import WelcomeAditya from '../assets/icons/Welcome.png';
	import MusicBurn from '../assets/icons/MusicBurn.png';
	import Play from '../assets/icons/play.png';
	import Tile1 from '../assets/icons/tile1.png';
	import Tile2 from '../assets/icons/tile2.png';
	import Tile3 from '../assets/icons/tile3.png';
	import Tile4 from '../assets/icons/tile4.png';
	import marketPlaceIcon from '../assets/icons/marketPlaceIcon.png';
	import DownArrowIcon from '../assets/icons/arrow-down.svg';
	import Player from '../lib/components/player/Player.svelte';
	import { navigate } from 'svelte-routing';
	import { isAuthenticated } from '../lib/middleware/auth';
	import { onMount } from 'svelte';

	onMount(() => {
		if (!isAuthenticated()) {
			navigate('/');
		}
	});

	const {
		elements: { root, content, trigger },
		states: { open },
	} = createCollapsible({ forceVisible: true, defaultOpen: false });
</script>

<div
	use:melt="{$root}"
	class="flex bg-[#121212] overflow-x-hidden w-full flex-col"
>
	<div
		class="w-full h-fit bg-[#121212] border-b border-[#3C3C4399] flex sticky top-0 items-center gap-8 py-2 px-8"
	>
		<div class="py-2 border-b border-[#3C3C4399]">
			<img src="{UnmuteLogo}" alt="Unmute Logo" class="max-w-120" />
		</div>

		<div class="w-1/2 mx-auto">
			<Player
				artists="Tame Impala"
				title="The lesser I know the better"
			/>
		</div>

		<div class="flex gap-4 ml-auto">
			<div class="absolute z-20 right-5 offset-top">
				{#if $open}
					<div
						use:melt="{$content}"
						transition:slide
						class="p-4 text-white rounded bg-brightGrey"
					>
						<button
							on:click="{signOut}"
							class="px-4 py-2 font-bold uppercase rounded bg-partyPurple text-spaceGrey"
						>
							Sign Out
						</button>
					</div>
				{/if}
			</div>
			<img
				src="https://i.pravatar.cc/300?img=3"
				alt="Avatar"
				class="w-10 h-10 rounded-full"
			/>

			<button use:melt="{$trigger}" aria-label="Toggle">
				{#if $open}
					<img
						src="{DownArrowIcon}"
						alt="Down Arrow icon"
						class="rotate-180"
					/>
				{:else}
					<img src="{DownArrowIcon}" alt="Down Arrow icon" />
				{/if}
			</button>
		</div>
	</div>
	<div class="flex flex-row">
		<div class=" bg-[#FFFFFF05] p-8 border-r border-[#3C3C4399]">
			<ul class="p-0 text-white list-none">
				<li class="mb-9">
					<Link to="/dashboard" class="flex items-center">
						<img
							src="{DashboardIcon}"
							alt="Dashboard"
							class="w-5 h-5 mr-2"
						/> Dashboard
					</Link>
				</li>
				<li class="mb-9">
					<Link to="/feedback" class="flex items-center">
						<img
							src="{DraftsIcon}"
							alt="Drafts"
							class="w-5 h-5 mr-2"
						/> Feedback
					</Link>
				</li>
				<li class="mb-9">
					<Link to="/chat" class="flex items-center">
						<img
							src="{ChatsIcon}"
							alt="Chats"
							class="w-5 h-5 mr-2"
						/> Chats
					</Link>
				</li>
				<li class="mb-9">
					<Link to="/explore" class="flex items-center">
						<img
							src="{ExploreIcon}"
							alt="Explore"
							class="w-5 h-5 mr-2"
						/> Explore
					</Link>
				</li>
				<li class="mb-9">
					<Link to="/marketplace" class="flex items-center">
						<img
							src="{marketPlaceIcon}"
							alt="Website Builder"
							class="w-5 h-5 mr-2"
						/> Market Place
					</Link>
				</li>
				<li class="mb-9">
					<a
						href="https://buy.stripe.com/test_4gw8zD5Jk6lg2088ww"
						class="flex items-center"
					>
						<img
							src="{PaymentsIcon}"
							alt="Payments"
							class="w-5 h-5 mr-2"
						/> Payments
					</a>
				</li>
				<li class="mb-9">
					<Link to="/contracts" class="flex items-center">
						<img
							src="{ContractsIcon}"
							alt="Contracts"
							class="w-5 h-5 mr-2"
						/> Contracts
					</Link>
				</li>
				<li class="mb-16">
					<Link to="/profile" class="flex items-center">
						<img
							src="{ProfileIcon}"
							alt="Profile"
							class="w-5 h-5 mr-2"
						/> Profile
					</Link>
				</li>
				<li class="mb-2">
					<Link to="/settings" class="flex items-center">
						<img
							src="{SettingsIcon}"
							alt="Settings"
							class="w-5 h-5 mr-2"
						/> Settings
					</Link>
				</li>
			</ul>
		</div>

		<div class="flex flex-col">
			<div class="p-8">
				<div class="flex items-center w-full">
					<img
						src="{WelcomeAditya}"
						alt="Welcome Aditya"
						class="mb-4 mr-4"
					/>
					<h2 class="mb-4 text-3xl font-semibold text-white">
						Welcome Aditya,
					</h2>
				</div>
				<p class="mb-4 ml-4 text-xl text-gray-300">
					Your Music Management Hub! Elevate your career with
					analytics, gigs, and fan connections. Let's make music
					magic!
				</p>
				<div class="flex flex-wrap">
					<img
						src="{RevenueGraph}"
						alt="Revenue Graph"
						class="h-56 mr-4 w-96"
					/>
					<img
						src="{AudienceGraph}"
						alt="Audience Graph"
						class="h-56 mr-4 w-96"
					/>
					<img src="{FanGraph}" alt="Fan Graph" class="h-56 w-96" />
				</div>
			</div>
			<div class="w-[80vw] p-8 overflow-auto no-scrollbar">
				<div class="flex items-center w-full">
					<img
						src="{MusicBurn}"
						alt="Welcome Aditya"
						class="mb-4 mr-4"
					/>
					<h2 class="mb-4 text-3xl font-semibold text-white">
						Your Music
					</h2>
				</div>
				<div class="flex w-screen gap-8 flex-nowrap">
					<div
						class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-[#252525] rounded-3xl p-4 scroll-snap-align-start"
					>
						<img
							src="{Tile1}"
							alt="Tile1"
							class="w-full h-auto mb-2"
						/>
						<h3 class="mb-2 text-lg font-semibold text-white">
							The Less I Know T...
						</h3>
						<button
							class="bg-[#785BEB] text-spaceGray w-full flex items-center justify-center px-4 py-2 rounded"
						>
							<img src="{Play}" alt="Play Button" class="mr-2" />
							Play
						</button>
					</div>
					<div
						class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-[#252525] rounded-3xl p-4 scroll-snap-align-start"
					>
						<img
							src="{Tile2}"
							alt="Tile2"
							class="w-full h-auto mb-2"
						/>
						<h3 class="mb-2 text-lg font-semibold text-white">
							Zakir
						</h3>
						<button
							class="bg-[#785BEB] text-spaceGray w-full flex items-center justify-center px-4 py-2 rounded"
						>
							<img src="{Play}" alt="Play Button" class="mr-2" />
							Play
						</button>
					</div>
					<div
						class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-[#252525] rounded-3xl p-4 scroll-snap-align-start"
					>
						<img
							src="{Tile3}"
							alt="Tile3"
							class="w-full h-auto mb-2"
						/>
						<h3 class="mb-2 text-lg font-semibold text-white">
							FE!N
						</h3>
						<button
							class="bg-[#785BEB] text-spaceGray w-full flex items-center justify-center px-4 py-2 rounded"
						>
							<img src="{Play}" alt="Play Button" class="mr-2" />
							Play
						</button>
					</div>
					<div
						class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-[#252525] rounded-3xl p-4 scroll-snap-align-start"
					>
						<img
							src="{Tile4}"
							alt="Tile4"
							class="w-full h-auto mb-2"
						/>
						<h3 class="mb-2 text-lg font-semibold text-white">
							Zakir
						</h3>
						<button
							class="bg-[#785BEB] text-spaceGray w-full flex items-center justify-center px-4 py-2 rounded"
						>
							<img src="{Play}" alt="Play Button" class="mr-2" />
							Play
						</button>
					</div>
					<div
						class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-[#252525] rounded-3xl p-4 scroll-snap-align-start"
					>
						<img
							src="{Tile4}"
							alt="Tile4"
							class="w-full h-auto mb-2"
						/>
						<h3 class="mb-2 text-lg font-semibold text-white">
							Zakir
						</h3>
						<button
							class="bg-[#785BEB] text-spaceGray w-full flex items-center justify-center px-4 py-2 rounded"
						>
							<img src="{Play}" alt="Play Button" class="mr-2" />
							Play
						</button>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	.no-scrollbar::-webkit-scrollbar {
		display: none;
	}
	.no-scrollbar {
		-ms-overflow-style: none;
		scrollbar-width: none;
	}

	.offset-top {
		top: calc(100% + 10px);
	}
</style>
