import classes from "./Header.module.css";
import {ActionIcon, Anchor, Image, Title, Tooltip} from "@mantine/core";
import TibiaWikiIcon from "~/assets/tibia-wiki.png";

interface DetailHeaderProps {
    item: string;
    wikiLink: string;
}

export function DetailHeader({ item, wikiLink }: DetailHeaderProps) {
    return (
        <header className={classes.header}>
            <Title className={classes.title}>{item}</Title>
            <Tooltip label="Go to TibiaWiki" openDelay={300}>
                <Anchor href={wikiLink} target="_blank">
                    <ActionIcon variant="default" aria-label="Tibia Wiki">
                        <Image src={TibiaWikiIcon as string} alt="Tibia Wiki" h={20} w={20}/>
                    </ActionIcon>
                </Anchor>
            </Tooltip>
        </header>
    )
}