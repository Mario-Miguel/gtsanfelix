/*eslint-disable*/
import React from "react";
import {Flex, Text} from "@chakra-ui/react";

export default function Footer() {
    // const linkTeal = useColorModeValue("teal.400", "red.200");=
    return (
        <Flex
            className={"foooter"}
            flexDirection={{base: "column", xl: "row",}}
            alignItems={{base: "center", xl: "start",}}
            justifyContent="space-between"
            px="10px"
            pb="20px"
        >
            <Text
                color="gray.400"
                textAlign={{base: "center", xl: "start"}}

            >
                &copy; {1900 + new Date().getYear()},{" "}
                <Text as="span">G.T. San Felix</Text>
            </Text>

        </Flex>
    );
}
