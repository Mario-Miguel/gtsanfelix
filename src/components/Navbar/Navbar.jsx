import {
	Box,
	Collapse,
	Flex,
	IconButton,
	Text,
	useColorModeValue,
	useBreakpointValue,
	useDisclosure,
} from "@chakra-ui/react";
import { FaBars, FaWindowClose } from "react-icons/fa";
import DesktopNav from "./DesktopNav";
import MobileNav from "./MobileNav";
import routes from "../../routes.jsx";

const Navbar = ({ navBarHeight }) => {
	const { isOpen, onToggle } = useDisclosure();

	return (
		<Box position={"fixed"} borderTop="0px" width={"100%"}>
			<Flex
				width={"100%"}
				bg={useColorModeValue("white", "gray.800")}
				color={useColorModeValue("gray.600", "white")}
				minH={navBarHeight}
				py={{ base: 2 }}
				px={{ base: 4 }}
				borderBottom={1}
				borderStyle={"solid"}
				borderColor={useColorModeValue("gray.200", "gray.900")}
				align={"center"}
			>
				<Flex flex={{ base: 1, md: "auto" }} ml={{ base: -2 }} display={{ base: "flex", md: "none" }}>
					<IconButton
						onClick={onToggle}
						icon={isOpen ? <FaWindowClose w={3} h={3} /> : <FaBars w={5} h={5} />}
						variant={"ghost"}
						aria-label={"Toggle Navigation"}
					/>
				</Flex>
				<Flex flex={{ base: 1 }} justify={{ base: "center", md: "start" }}>
					<Text
						textAlign={useBreakpointValue({ base: "center", md: "left" })}
						fontFamily={"heading"}
						color={useColorModeValue("gray.800", "white")}
					>
						Logo
					</Text>

					<Flex display={{ base: "none", md: "flex" }} ml={10}>
						<DesktopNav navItems={routes} />
					</Flex>
				</Flex>
			</Flex>

			<Collapse in={isOpen} animateOpacity>
				<MobileNav navItems={routes} />
			</Collapse>
		</Box>
	);
};

export default Navbar;
