// Chakra imports
import { ChakraProvider, Portal, useDisclosure, Text } from "@chakra-ui/react";
// Layout components
import React from "react";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";
// Custom Chakra theme
import theme from "../theme/theme.jsx";
import PanelContent from "../components/Layout/PanelContent.jsx";
import PanelContainer from "../components/Layout/PanelContainer.jsx";
import Navbar from "../components/Navbar/Navbar.jsx";
import MainPanel from "../components/Layout/MainPanel.jsx";
import Footer from "../components/Footer/Footer.jsx";
import { Outlet } from "react-router-dom";

// Custom components

export default function UserLayout() {
	const { isOpen, onOpen, onClose } = useDisclosure();
	const navBarHeight = "100px";

	return (
		<ChakraProvider theme={theme} resetCss={false}>
			<MainPanel w={"100%"}>
				<Portal>
					<Navbar navBarHeight={navBarHeight} />
				</Portal>

				<PanelContent marginTop={navBarHeight}>
					<PanelContainer>
						<Outlet />
					</PanelContainer>
				</PanelContent>

				<Footer />
			</MainPanel>
		</ChakraProvider>
	);
}
