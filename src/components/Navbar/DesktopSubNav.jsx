import { Box, Flex, Icon, Stack, Text, useColorModeValue } from "@chakra-ui/react";
import { FaChevronDown } from "react-icons/fa";
import { Link } from "react-router-dom";

const DesktopSubNav = ({ name, path }) => {
	return (
		<Link to={path} p={2} rounded={"md"} _hover={{ bg: useColorModeValue("pink.50", "gray.900") }}>
			<Stack direction={"row"} align={"center"}>
				<Box>
					<Text transition={"all .3s ease"} _groupHover={{ color: "pink.400" }} fontWeight={500}>
						{name}
					</Text>
				</Box>
				<Flex
					transition={"all .3s ease"}
					transform={"translateX(-10px)"}
					opacity={0}
					_groupHover={{ opacity: "100%", transform: "translateX(0)" }}
					justify={"flex-end"}
					align={"center"}
					flex={1}
				>
					<Icon color={"pink.400"} w={5} h={5} as={FaChevronDown} />
				</Flex>
			</Stack>
		</Link>
	);
};

export default DesktopSubNav;
