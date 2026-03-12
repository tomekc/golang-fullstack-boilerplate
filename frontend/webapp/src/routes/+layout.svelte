<script lang="ts">
    import '$lib/styles/admin-one.css';
    import favicon from '$lib/assets/favicon.svg';
    import { page } from '$app/state';
    import { AdminNavbar, NavbarDropdown, NavbarDropdownItem, NavbarLink,
             AdminSidebar, SidebarTools, SidebarMenu, SidebarMenuSection, SidebarMenuItem } from 'svebulma';

    let { children } = $props();
    const pathname = $derived(page.url.pathname);
</script>

<svelte:head>
    <link rel="icon" href={favicon}/>
</svelte:head>

<AdminNavbar>
    {#snippet navbarEnd()}
        <NavbarDropdown label="John Doe" avatarUrl="https://avatars.dicebear.com/v2/initials/john-doe.svg" hasDivider>
            <NavbarDropdownItem icon="account" label="My Profile" />
            <NavbarDropdownItem icon="settings" label="Settings" />
            <NavbarDropdownItem divider />
            <NavbarDropdownItem icon="logout" label="Log out" />
        </NavbarDropdown>
        <NavbarLink icon="help-circle-outline" label="About" href="/about" hasDivider desktopIconOnly />
        <NavbarLink icon="logout" label="Log out" desktopIconOnly />
    {/snippet}
</AdminNavbar>

<AdminSidebar>
    <SidebarTools>
        <span><b>Admin</b> One</span>
    </SidebarTools>
    <SidebarMenu>
        <SidebarMenuSection label="General">
            <SidebarMenuItem label="Dashboard" href="/" icon="desktop-mac" active={pathname === '/'} />
        </SidebarMenuSection>
        <SidebarMenuSection label="Examples">
            <SidebarMenuItem label="Hello API" href="/hello" icon="clock-outline" active={pathname === '/hello'} />
            <SidebarMenuItem label="About" href="/about" icon="help-circle" active={pathname === '/about'} />
        </SidebarMenuSection>
    </SidebarMenu>
</AdminSidebar>

{@render children()}

<footer class="footer">
    <div class="container-fluid">
        <div class="level">
            <div class="level-left">
                <div class="level-item">
                    © 2025, Golang-SvelteKit Boilerplate
                </div>
            </div>
        </div>
    </div>
</footer>
