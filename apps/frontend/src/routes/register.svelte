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
				navigate('/dashboard');
			})
			.catch((err) => console.error(err));
	};

	onMount(() => {
		if (isAuthenticated()) {
			navigate('/dashboard');
		}
	});
</script>

<div
	class="flex flex-col items-center justify-center w-full min-h-screen gap-8 p-8 bg-spaceGrey"
>
	<Link to="/">
		<img
			src="{UnmuteLogo}"
			alt="Unmute Logo"
			class="absolute top-8 left-8"
		/>
	</Link>
	<div class="w-full text-white sm:w-2/3 lg:w-1/3">
		<Link to="/login" class="uppercase text-darkWhite">Login</Link><span
			class="mx-3 border-l"
		></span><span class="uppercase text-partyPurple">Sign Up</span>
	</div>

	<div class="flex flex-col w-full gap-2 sm:w-2/3 lg:w-1/3">
		<h1 class="text-2xl tracking-wider text-white font-ownersTrialWide">
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
			class="mb-2 text-sm font-semibold text-white uppercase"
		>
			Full Name
		</label>
		<input
			type="text"
			name="fullname"
			id="fullname"
			placeholder="Ramesh Waterwala"
			class="p-2 text-white bg-transparent border rounded border-brightGrey active:border-partyPurple active:ring-partyPurple"
		/>

		<label
			for="username"
			class="mt-8 mb-2 text-sm font-semibold text-white uppercase"
		>
			Username
		</label>
		<input
			type="text"
			name="username"
			id="username"
			placeholder="johndoe"
			class="p-2 text-white bg-transparent border rounded border-brightGrey active:border-partyPurple active:ring-partyPurple"
		/>

		<label
			for="email"
			class="mt-8 mb-2 text-sm font-semibold text-white uppercase"
		>
			Email
		</label>
		<input
			type="email"
			name="email"
			id="email"
			placeholder="example@gmail.com"
			class="p-2 text-white bg-transparent border rounded border-brightGrey active:border-partyPurple active:ring-partyPurple"
		/>

		<label
			for="password"
			class="mt-8 mb-2 text-sm font-semibold text-white uppercase"
		>
			Password
		</label>
		<input
			type="password"
			name="password"
			id="password"
			placeholder="******"
			class="p-2 text-white bg-transparent border rounded border-brightGrey active:border-partyPurple active:ring-partyPurple"
		/>

		<label
			for="confpassword"
			class="mt-8 mb-2 text-sm font-semibold text-white uppercase"
		>
			Confirm Password
		</label>
		<input
			type="password"
			name="confpassword"
			id="confpassword"
			placeholder="******"
			class="p-2 mb-8 text-white bg-transparent border rounded border-brightGrey active:border-partyPurple active:ring-partyPurple"
		/>

		<button
			type="submit"
			class="py-2 font-bold uppercase rounded bg-partyPurple text-spaceGrey"
		>
			Sign Up
		</button>
	</form>

	<button
		class="flex items-center justify-center w-full gap-2 py-2 font-bold uppercase bg-transparent border rounded border-partyPurple text-darkWhite hover:bg-partyPurple sm:w-1/2 lg:w-1/3"
	>
		<img src="{BrandGoogle}" alt="Google Icon" class="h-4" />
		<span> Sign in with Google </span>
	</button>

	{#if errorText}
		<p class="text-white">{errorText}</p>
	{/if}
</div>
