<script lang="ts">
	import { Link, navigate } from 'svelte-routing';
	import Cookies from 'js-cookie';
	import { onMount } from 'svelte';
	import { isAuthenticated } from '../lib/middleware/auth';

	import UnmuteLogo from '../assets/logo.svg';
	import BrandGoogle from '../assets/icons/brand-google.svg';

	let errorText: string | undefined;

	type TCredentials = {
		fullname?: string;
		username?: string;
		email?: string;
		password?: string;
		role?: string;
	};
	const onSumbit = (e: SubmitEvent) => {
		errorText = undefined;

		const formData = new FormData(
			(e.target as HTMLFormElement) ?? undefined
		);

		const credentials: TCredentials = {};
		for (let [key, value] of formData) {
			credentials[key.toString() as keyof TCredentials] =
				value.toString();
		}
		credentials.role = 'artist';

		console.log(credentials);

		fetch(`${import.meta.env.VITE_SERVER_URL}/api/signup`, {
			method: 'POST',
			cache: 'no-cache',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(credentials),
		})
			.then((res) => {
				return res.json();
			})
			.then((data) => {
				if (data.error) {
					errorText = data.error;
					return;
				}

				Cookies.set('PodiDosa', data.token);
				navigate('/protected');
			})
			.catch((err) => console.error(err));
	};

	onMount(() => {
		if (isAuthenticated()) {
			navigate('/protected');
		}
	});
</script>

<div
	class="w-full min-h-screen bg-spaceGrey flex flex-col items-center justify-center gap-8 p-8"
>
	<img
		src="{UnmuteLogo}"
		alt="Unmute Logo"
		class="absolute top-8 md:left-8 right-8"
	/>
	<div class="w-full sm:w-2/3 lg:w-1/3 text-white">
		<Link to="/login" class="uppercase text-darkWhite">Login</Link><span
			class="border-l mx-3"
		></span><span class="uppercase text-partyPurple">Sign Up</span>
	</div>

	<div class="flex flex-col gap-2 w-full sm:w-2/3 lg:w-1/3">
		<h1 class="text-white text-2xl font-ownersTrialWide tracking-wider">
			Welcome,
		</h1>
		<p class="text-darkWhite">Please enter your details below to sign up</p>
	</div>

	<form
		on:submit|preventDefault="{onSumbit}"
		class="flex flex-col w-full sm:w-2/3 lg:w-1/3"
	>
		<label
			for="fullname"
			class="text-white text-sm mb-2 uppercase font-semibold"
		>
			Full Name
		</label>
		<input
			type="text"
			name="fullname"
			id="fullname"
			placeholder="Ramesh Waterwala"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2 text-white"
		/>

		<label
			for="username"
			class="text-white text-sm mt-8 mb-2 uppercase font-semibold"
		>
			Username
		</label>
		<input
			type="text"
			name="username"
			id="username"
			placeholder="johndoe"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2 text-white"
		/>

		<label
			for="email"
			class="text-white text-sm mt-8 mb-2 uppercase font-semibold"
		>
			Email
		</label>
		<input
			type="email"
			name="email"
			id="email"
			placeholder="example@gmail.com"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2 text-white"
		/>

		<label
			for="password"
			class="text-white text-sm mt-8 mb-2 uppercase font-semibold"
		>
			Password
		</label>
		<input
			type="password"
			name="password"
			id="password"
			placeholder="******"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2 text-white"
		/>

		<label
			for="confpassword"
			class="text-white text-sm mt-8 mb-2 uppercase font-semibold"
		>
			Confirm Password
		</label>
		<input
			type="password"
			name="confpassword"
			id="confpassword"
			placeholder="******"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2 mb-8 text-white"
		/>

		<button
			type="submit"
			class="bg-partyPurple text-spaceGrey rounded py-2 font-bold uppercase"
		>
			Sign Up
		</button>
	</form>

	<button
		class="flex justify-center items-center gap-2 border border-partyPurple text-darkWhite bg-transparent hover:bg-partyPurple w-full sm:w-1/2 lg:w-1/3 rounded py-2 font-bold uppercase"
	>
		<img src="{BrandGoogle}" alt="Google Icon" class="h-4" />
		<span> Sign in with Google </span>
	</button>

	{#if errorText}
		<p class="text-white">{errorText}</p>
	{/if}
</div>
