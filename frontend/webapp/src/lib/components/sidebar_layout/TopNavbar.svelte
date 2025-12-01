<script lang="ts">
    import { Navbar, NavbarBrand } from '@sveltestrap/sveltestrap';

    interface Props {
        appName?: string;
        version?: string;
    }

    let { appName = 'Boilerplate', version = '0.0.1' }: Props = $props();
</script>

<Navbar dark color="dark" expand class="sb-topnav px-3">
    <NavbarBrand href="/">{appName}</NavbarBrand>
    <span class="ms-auto text-white-50 small">version {version}</span>
</Navbar>

<style>
    :global(.sb-topnav) {
        z-index: 1039;
        height: 56px;
    }
</style>
