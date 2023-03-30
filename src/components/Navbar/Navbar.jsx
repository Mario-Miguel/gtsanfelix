// Chakra Imports
import {Flex,} from "@chakra-ui/react";
import React from "react";


export default function Navbar({navBarHeight='100px'}) {

    // Here are all the props that may change depending on navbar's type or state.(secondary, variant, scrolled)
    let navbarPosition = "absolute";
    let navbarFilter = "none";
    let navbarBackdrop = "blur(21px)";
    let navbarShadow = "xl";
    let navbarBg = "none";
    let navbarBorder = "transparent";
    let secondaryMargin = "0px";
    let paddingX = "0px";

    return (
            <Flex
                id={"navbar"}
                position={navbarPosition}
                boxShadow={navbarShadow}
                bg={navbarBg}
                borderColor={navbarBorder}
                filter={navbarFilter}
                backdropFilter={navbarBackdrop}
                borderWidth="1.5px"
                borderStyle="solid"
                alignItems={{xl: "center"}}
                display="flex"
                minH={navBarHeight}
                justifyContent={{xl: "center"}}
                lineHeight="25.6px"
                mx="auto"
                mt={secondaryMargin}
                pb="8px"
                right={"0px"}
                px={{sm: paddingX, md: "30px",}}
                pt="8px"
                top="0px"
                w={"100%"}
            >
                <Flex w="100%" h={"100%"} flexDirection={{sm: "column", md: "row",}} alignItems={{xl: "center"}}>

                </Flex>
            </Flex>
    );
}
