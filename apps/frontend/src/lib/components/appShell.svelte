<script lang="ts">
	import { Link } from 'svelte-routing';
	import UnmuteIcon from '../../assets/icons/unmute_icon.png';
	import DashboardIcon from '../../assets/icons/dashboard.png';
	import DraftsIcon from '../../assets/icons/drafts.png';
	import ChatsIcon from '../../assets/icons/chats.png';
	import ExploreIcon from '../../assets/icons/explore.png';
	import PaymentsIcon from '../../assets/icons/payments.png';
	import ContractsIcon from '../../assets/icons/contracts.png';
	import ProfileIcon from '../../assets/icons/profile.png';
	import SettingsIcon from '../../assets/icons/settings.png';
	import marketPlaceIcon from '../../assets/icons/marketPlaceIcon.png';
</script>

<div class="w-full min-h-screen bg-[#121212] flex flex-col">
	<div class="pl-6 pt-6 pb-4 border-b border-[#3C3C4399]">
		<img src="{UnmuteIcon}" alt="Unmute Logo" class="w-8 h-8 mr-2" />
	</div>

	<div class="flex flex-grow w-full">
		<div
			class="w-fit bg-[#FFFFFF05] px-8 py-12 border-r border-[#3C3C4399]"
		>
			<ul class="p-0 list-none text-darkWhite">
				<li class="mb-10">
					<Link to="/dashboard" class="flex items-center">
						<img
							src="{DashboardIcon}"
							alt="Dashboard"
							class="w-5 h-5 mr-2"
						/>
					</Link>
				</li>
				<li class="mb-10">
					<Link to="/drafts" class="flex items-center">
						<img
							src="{DraftsIcon}"
							alt="Drafts"
							class="w-5 h-5 mr-2"
						/>
					</Link>
				</li>
				<li class="mb-10">
					<Link to="/chat" class="flex items-center">
						<img
							src="{ChatsIcon}"
							alt="Chats"
							class="w-5 h-5 mr-2"
						/>
					</Link>
				</li>
				<li class="mb-10">
					<Link to="/explore" class="flex items-center">
						<img
							src="{ExploreIcon}"
							alt="Explore"
							class="w-5 h-5 mr-2"
						/>
					</Link>
				</li>
				<li class="mb-10">
					<Link to="/marketplace" class="flex items-center">
						<img
							src="{marketPlaceIcon}"
							alt="MarketPlace"
							class="w-5 h-5 mr-2"
						/>
					</Link>
				</li>
				<li class="mb-10">
					<Link to="/payments" class="flex items-center">
						<img
							src="{PaymentsIcon}"
							alt="Payments"
							class="w-5 h-5 mr-2"
						/>
					</Link>
				</li>
				<li class="mb-10">
					<Link to="/contracts" class="flex items-center">
						<img
							src="{ContractsIcon}"
							alt="Contracts"
							class="w-5 h-5 mr-2"
						/>
					</Link>
				</li>
				<li class="mb-16">
					<Link to="/profile" class="flex items-center">
						<img
							src="{ProfileIcon}"
							alt="Profile"
							class="w-5 h-5 mr-2"
						/>
					</Link>
				</li>
				<li class="mb-2">
					<Link to="/settings" class="flex items-center">
						<img
							src="{SettingsIcon}"
							alt="Settings"
							class="w-5 h-5 mr-2"
						/>
					</Link>
				</li>
			</ul>
		</div>

		<div
			class="w-full p-4"
			style="background: linear-gradient(180deg, rgba(20, 179, 118, 0.05) 0%, rgba(17, 25, 40, 0.00) 50%);"
		>
			<slot />
		</div>
	</div>
</div>
