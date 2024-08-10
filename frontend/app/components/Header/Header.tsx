import classes from "./Header.module.css";
import {Title} from "@mantine/core";
import {ColorSchemeToggle} from "~/components/ColorSchemeToggle/ColorSchemeToggle";

export function Header() {
    return (
        <header className={classes.header}>
            <Title className={classes.title}>Tibia Mkt</Title>
            <ColorSchemeToggle/>
        </header>
    )
}