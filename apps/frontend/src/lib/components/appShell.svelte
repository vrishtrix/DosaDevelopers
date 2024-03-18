<script lang="ts">
    import { Link } from 'svelte-routing';
    import Dashboard from '../../routes/dashboard.svelte';
    import Drafts from '../../routes/drafts.svelte';
    import Chats from '../../routes/chats.svelte';
    import Explore from '../../routes/explore.svelte';
    import WebsiteBuilder from '../../routes/websiteBuilder.svelte';
    import Payments from '../../routes/payments.svelte';
    import Contracts from '../../routes/contracts.svelte';
    import Profile from '../../routes/profile.svelte';
    import Settings from '../../routes/settings.svelte';

    import UnmuteIcon from '../assets/icons/unmute_icon.png';
    import DashboardIcon from '../assets/icons/dashboard.png'; 
    import DraftsIcon from '../assets/icons/drafts.png'; 
    import ChatsIcon from '../assets/icons/chats.png'; 
    import ExploreIcon from '../assets/icons/explore.png'; 
    import WebsiteBuilderIcon from '../assets/icons/websiteBuilder.png'; 
    import PaymentsIcon from '../assets/icons/payments.png'; 
    import ContractsIcon from '../assets/icons/contracts.png'; 
    import ProfileIcon from '../assets/icons/profile.png'; 
    import SettingsIcon from '../assets/icons/settings.png'; 
</script>

<div class="w-full min-h-screen bg-[#121212] flex flex-col">
    <div class="pl-6 pt-6 pb-4 border-b border-[#3C3C4399]">
        <img src={UnmuteIcon} alt="Unmute Logo" class="w-8 h-8 mr-2" />
    </div>

    <div class="flex-grow flex">
        <div class="w-fit bg-[#FFFFFF05] px-8 py-12 border-r border-[#3C3C4399]">
            <ul class="list-none p-0 text-darkWhite">
                <li class="mb-10">
                    <Link to="/dashboard" class="flex items-center">
                        <img src={DashboardIcon} alt="Dashboard" class="w-5 h-5 mr-2" />
                    </Link>
                </li>
                <li class="mb-10">
                    <Link to="/drafts" class="flex items-center">
                        <img src={DraftsIcon} alt="Drafts" class="w-5 h-5 mr-2" /> 
                    </Link>
                </li>
                <li class="mb-10">
                    <Link to="/chats" class="flex items-center">
                        <img src={ChatsIcon} alt="Chats" class="w-5 h-5 mr-2" /> 
                    </Link>
                </li>
                <li class="mb-10">
                    <Link to="/explore" class="flex items-center">
                        <img src={ExploreIcon} alt="Explore" class="w-5 h-5 mr-2" /> 
                    </Link>
                </li>
                <li class="mb-10">
                    <Link to="/website-builder" class="flex items-center">
                        <img src={WebsiteBuilderIcon} alt="Website Builder" class="w-5 h-5 mr-2" /> 
                    </Link>
                </li>
                <li class="mb-10">
                    <Link to="/payments" class="flex items-center">
                        <img src={PaymentsIcon} alt="Payments" class="w-5 h-5 mr-2" /> 
                    </Link>
                </li>
                <li class="mb-10">
                    <Link to="/contracts" class="flex items-center">
                        <img src={ContractsIcon} alt="Contracts" class="w-5 h-5 mr-2" /> 
                    </Link>
                </li>
                <li class="mb-16">
                    <Link to="/profile" class="flex items-center">
                        <img src={ProfileIcon} alt="Profile" class="w-5 h-5 mr-2" /> 
                    </Link>
                </li>
                <li class="mb-2">
                    <Link to="/settings" class="flex items-center">
                        <img src={SettingsIcon} alt="Settings" class="w-5 h-5 mr-2" /> 
                    </Link>
                </li>
            </ul>
        </div>
    </div>
</div>
