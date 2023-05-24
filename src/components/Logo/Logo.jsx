import { Box, Text } from "@chakra-ui/react";

const Logo = (props) => {
	return (
		<Box {...props}>
			<Text fontSize="lg" fontWeight="bold" color={"black"}>
				Logo
			</Text>
		</Box>
	);
};

export default Logo;
