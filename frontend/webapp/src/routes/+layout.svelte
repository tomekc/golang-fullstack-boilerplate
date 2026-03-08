<script lang="ts">
    import favicon from '$lib/assets/favicon.svg';
    import TopNavbar from '$lib/components/sidebar_layout/TopNavbar.svelte';
    import Sidebar from '$lib/components/sidebar_layout/Sidebar.svelte';
    import MainContent from '$lib/components/sidebar_layout/MainContent.svelte';
    import SidebarSection from '$lib/components/sidebar_layout/SidebarSection.svelte';
    import SidebarItem from '$lib/components/sidebar_layout/SidebarItem.svelte';

    let {children} = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon}/>
</svelte:head>

<div class="is-flex is-flex-direction-column" style="min-height: 100vh;">
    <TopNavbar appName="Boilerplate" version="0.0.1" />
    <div id="layoutSidenav" class="is-flex is-flex-grow-1">
        <Sidebar userInfo="Admin User">
            <SidebarSection title="Core">
                <SidebarItem href="/" icon="house">Home</SidebarItem>
                <SidebarItem href="/hello" icon="clock">Hello, world!</SidebarItem>
                <SidebarItem href="/about" icon="address-card">About</SidebarItem>
            </SidebarSection>
        </Sidebar>

        <MainContent copyright="Golang-SvelteKit Boilerplate 2025">
            {@render children()}
        </MainContent>
    </div>
</div>

<style>
    #layoutSidenav {
        display: flex;
    }
</style>
