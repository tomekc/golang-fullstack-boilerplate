<script lang="ts">
    import favicon from '../lib/assets/favicon.svg';

    import {Container, Nav, Navbar, NavbarBrand, NavItem, NavLink} from '@sveltestrap/sveltestrap';

    let {children} = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon}/>
</svelte:head>

<Navbar color="dark" dark expand class="px-3">
    <NavbarBrand href="/">Boilerplate</NavbarBrand>
    <Nav navbar pills class="me-auto">
        <NavItem>
            <NavLink href="/">Home</NavLink>
        </NavItem>
        <NavItem>
            <NavLink href="/about">About</NavLink>
        </NavItem>
    </Nav>
</Navbar>
<Container fluid class="pt-2">
    {@render children()}
</Container>
