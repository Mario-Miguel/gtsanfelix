import { Collapse, Flex, Link, Stack, Text, Icon, useDisclosure, useColorModeValue } from "@chakra-ui/react";
import { FaChevronDown, FaChevronRight } from "react-icons/fa";

const MobileNavItem = ({ label, name, path, icon, children }) => {
	const { isOpen, onToggle } = useDisclosure();

	return (
		<Stack spacing={4} onClick={children && onToggle}>
			<Flex
				py={2}
				as={Link}
				href={path ?? "#"}
				justify={"space-between"}
				align={"center"}
				_hover={{
					textDecoration: "none",
				}}
			>
				<Text fontWeight={600} color={useColorModeValue("gray.600", "gray.200")}>
					{name}
				</Text>
				{children && (
					<Icon
						as={FaChevronRight}
						transition={"all .25s ease-in-out"}
						transform={isOpen ? "rotate(180deg)" : ""}
						w={6}
						h={6}
					/>
				)}
			</Flex>

			<Collapse in={isOpen} animateOpacity style={{ marginTop: "0!important" }}>
				<Stack
					mt={2}
					pl={4}
					borderLeft={1}
					borderStyle={"solid"}
					borderColor={useColorModeValue("gray.200", "gray.700")}
					align={"start"}
				>
					{children &&
						children.map((child) => (
							<Link key={child.label} py={2} href={child.path}>
								{child.name}
							</Link>
						))}
				</Stack>
			</Collapse>
		</Stack>
	);
};

export default MobileNavItem;
