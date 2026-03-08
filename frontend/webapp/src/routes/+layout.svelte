<script lang="ts">
    import favicon from '$lib/assets/favicon.svg';
    import { page } from '$app/state';
    import {
        TopNavbar,
        Sidebar,
        SidebarSection,
        SidebarItem,
        MainContent,
        PageFooter,
    } from 'svebulma';

    let { children } = $props();
    const pathname = $derived(page.url.pathname);
</script>

<svelte:head>
    <link rel="icon" href={favicon}/>
</svelte:head>

<div class="is-flex is-flex-direction-column" style="min-height: 100vh;">
    <TopNavbar appName="Boilerplate" version="0.0.1" />
    <div class="is-flex is-flex-grow-1">
        <Sidebar userInfo="Admin User">
            <SidebarSection title="Core">
                <SidebarItem href="/" icon="house" active={pathname === '/'}>Home</SidebarItem>
                <SidebarItem href="/hello" icon="clock" active={pathname === '/hello'}>Hello, world!</SidebarItem>
                <SidebarItem href="/about" icon="address-card" active={pathname === '/about'}>About</SidebarItem>
            </SidebarSection>
        </Sidebar>
        <MainContent>
            {@render children()}
        </MainContent>
    </div>
    <PageFooter copyright="Golang-SvelteKit Boilerplate 2025" />
</div>
