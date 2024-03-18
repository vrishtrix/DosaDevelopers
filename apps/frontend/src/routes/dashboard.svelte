<script lang="ts">
    import { Link } from 'svelte-routing';
    import Dashboard from './dashboard.svelte';
    import Drafts from './drafts.svelte';
    import Chats from './chats.svelte';
    import Explore from './explore.svelte';
    import WebsiteBuilder from './websiteBuilder.svelte';
    import Payments from './payments.svelte';
    import Contracts from './contracts.svelte';
    import Profile from './profile.svelte';
    import Settings from './settings.svelte';

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
    import MusicBurn from '../assets/icons/MusicBurn.png'
    import Play from '../assets/icons/play.png'; 
    import Tile1 from '../assets/icons/tile1.png'; 
</script>
<style>
    .no-scrollbar::-webkit-scrollbar {
        display: none;
        }
    .no-scrollbar {
        -ms-overflow-style: none;  /* IE and Edge */
        scrollbar-width: none;  /* Firefox */
    }
</style>

<div class="flex bg-[#121212] overflow-x-hidden w-full flex-col">
<div class="w-full h-fit bg-black flex">
    <div class="pl-8 pt-4 pb-2 border-b border-[#3C3C4399]">
        <img src={UnmuteLogo} alt="Unmute Logo" class="max-w-120" />
    </div>
</div>
    <div class = "flex flex-row">
        <div class=" bg-[#FFFFFF05] p-8 border-r border-[#3C3C4399]">
            <ul class="list-none p-0 text-white">
                
                <li class="mb-9">
                    <Link to="/dashboard" class="flex items-center">
                        <img src={DashboardIcon} alt="Dashboard" class="w-5 h-5 mr-2" /> Dashboard
                    </Link>
                </li>
                <li class="mb-9">
                    <Link to="/drafts" class="flex items-center">
                        <img src={DraftsIcon} alt="Drafts" class="w-5 h-5 mr-2" /> Drafts
                    </Link>
                </li>
                <li class="mb-9">
                    <Link to="/chats" class="flex items-center">
                        <img src={ChatsIcon} alt="Chats" class="w-5 h-5 mr-2" /> Chats
                    </Link>
                </li>
                <li class="mb-9">
                    <Link to="/explore" class="flex items-center">
                        <img src={ExploreIcon} alt="Explore" class="w-5 h-5 mr-2" /> Explore
                    </Link>
                </li>
                <li class="mb-9">
                    <Link to="/website-builder" class="flex items-center">
                        <img src={WebsiteBuilderIcon} alt="Website Builder" class="w-5 h-5 mr-2" /> Website Builder
                    </Link>
                </li>
                <li class="mb-9">
                    <Link to="/payments" class="flex items-center">
                        <img src={PaymentsIcon} alt="Payments" class="w-5 h-5 mr-2" /> Payments
                    </Link>
                </li>
                <li class="mb-9">
                    <Link to="/contracts" class="flex items-center">
                        <img src={ContractsIcon} alt="Contracts" class="w-5 h-5 mr-2" /> Contracts
                    </Link>
                </li>
                <li class="mb-16">
                    <Link to="/profile" class="flex items-center">
                        <img src={ProfileIcon} alt="Profile" class="w-5 h-5 mr-2" /> Profile
                    </Link>
                </li>
                <li class="mb-2">
                    <Link to="/settings" class="flex items-center">
                        <img src={SettingsIcon} alt="Settings" class="w-5 h-5 mr-2" /> Settings
                    </Link>
                </li>
            </ul>
        </div>
        
        <div class="flex flex-col">
            <div class="p-8">
                <div class="flex w-full items-center">
                    <img src="{WelcomeAditya}" alt="Welcome Aditya" class="mr-4 mb-4">
                    <h2 class="text-white text-3xl font-semibold mb-4">Welcome Aditya,</h2>
                </div>
                <p class="text-gray-300 text-xl mb-4 ml-4">Your Music Management Hub! Elevate your career with analytics, gigs, and fan connections. Let's make music magic!</p>
                <div class="flex flex-wrap">
                    <img src="{RevenueGraph}" alt="Revenue Graph" class="w-96 h-56 mr-4">
                    <img src="{AudienceGraph}" alt="Audience Graph" class="w-96 h-56 mr-4">
                    <img src="{FanGraph}" alt="Fan Graph" class="w-96 h-56">
                </div>
            </div>
            <div class="w-[80vw] p-8 overflow-auto no-scrollbar">
                <div class="flex items-center w-full">
                    <img src="{ MusicBurn }" alt="Welcome Aditya" class="mr-4 mb-4">
                    <h2 class="text-white text-3xl font-semibold mb-4">Your Music</h2>
                </div>
                <div class="flex flex-nowrap w-screen gap-8">
                    <div class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-gray-800 rounded-lg p-4 scroll-snap-align-start">
                        <img src="{ Tile1 }" alt="Tile1" class="w-full h-auto mb-2">
                        <h3 class="text-white text-lg font-semibold mb-2">Title 1</h3>
                        <button class="bg-blue-500 text-white px-4 py-2 rounded">Play</button>
                    </div>
                    <div class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-gray-800 rounded-lg p-4 scroll-snap-align-start">
                        <img src="{ Tile1 }" alt="Tile1" class="w-full h-auto mb-2">
                        <h3 class="text-white text-lg font-semibold mb-2">Title 1</h3>
                        <button class="bg-blue-500 text-white px-4 py-2 rounded">Play</button>
                    </div>
                    <div class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-gray-800 rounded-lg p-4 scroll-snap-align-start">
                        <img src="{ Tile1 }" alt="Tile1" class="w-full h-auto mb-2">
                        <h3 class="text-white text-lg font-semibold mb-2">Title 1</h3>
                        <button class="bg-blue-500 text-white px-4 py-2 rounded">Play</button>
                    </div>
                    <div class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-gray-800 rounded-lg p-4 scroll-snap-align-start">
                        <img src="{ Tile1 }" alt="Tile1" class="w-full h-auto mb-2">
                        <h3 class="text-white text-lg font-semibold mb-2">Title 1</h3>
                        <button class="bg-blue-500 text-white px-4 py-2 rounded">Play</button>
                    </div>
                    <div class="w-64 flex-shrink-0 flex flex-col items-center justify-center bg-gray-800 rounded-lg p-4 scroll-snap-align-start">
                        <img src="{ Tile1 }" alt="Tile1" class="w-full h-auto mb-2">
                        <h3 class="text-white text-lg font-semibold mb-2">Title 1</h3>
                        <button class="bg-blue-500 text-white px-4 py-2 rounded">Play</button>
                    </div>
                </div>
                
            </div>
        </div>
        
    </div>

</div>
