import {
	Box,
	Button,
	Flex,
	FormControl,
	FormLabel,
	Heading,
	IconButton,
	Input,
	InputGroup,
	InputLeftElement,
	Link,
	Stack,
	Textarea,
	Tooltip,
	useClipboard,
	useColorModeValue,
	VStack,
} from "@chakra-ui/react";
import React from "react";
import { BsFacebook, BsInstagram, BsPerson } from "react-icons/bs";
import { MdEmail, MdOutlineEmail } from "react-icons/md";

const Contact = () => {
	const { hasCopied, onCopy } = useClipboard("gtsnafelix@gmail.com");

	return (
		<Flex bg={useColorModeValue("gray.100", "gray.900")} align="center" justify="center" id="contact">
			<Box borderRadius="lg" m={{ base: 5, md: 16, lg: 10 }}>
				<Box>
					<VStack spacing={{ base: 4, md: 8, lg: 20 }}>
						<Heading
							fontSize={{
								base: "4xl",
								md: "5xl",
							}}
						>
							¡Ponte en contautu con nos!
						</Heading>

						<Stack spacing={{ base: 4, md: 8, lg: 20 }} direction={{ base: "column", md: "row" }}>
							<Stack align="center" justify="space-around" direction={{ base: "row", md: "column" }}>
								<Tooltip
									label={hasCopied ? "¡Email Copiau!" : "Copiar Email"}
									closeOnClick={false}
									hasArrow
								>
									<IconButton
										aria-label="email"
										variant="ghost"
										size="lg"
										fontSize="3xl"
										icon={<MdEmail />}
										_hover={{
											bg: "blue.500",
											color: useColorModeValue("white", "gray.700"),
										}}
										onClick={onCopy}
										isRound
									/>
								</Tooltip>

								<Link
									href="https://www.instagram.com/gtsanfelixvaldesoto/"
									target="_blank"
									rel="noopener noreferrer"
								>
									<IconButton
										aria-label="instagram"
										variant="ghost"
										size="lg"
										fontSize="3xl"
										icon={<BsInstagram />}
										_hover={{
											bg: "blue.500",
											color: useColorModeValue("white", "gray.700"),
										}}
										isRound
									/>
								</Link>

								<Link
									href="https://www.facebook.com/g.t.sanfelix.valdesoto"
									target="_blank"
									rel="noopener noreferrer"
								>
									<IconButton
										aria-label="facebook"
										variant="ghost"
										size="lg"
										icon={<BsFacebook size={30} />}
										_hover={{
											bg: "blue.500",
											color: useColorModeValue("white", "gray.700"),
										}}
										isRound
									/>
								</Link>
							</Stack>

							<Box
								bg={useColorModeValue("white", "gray.700")}
								borderRadius="lg"
								p={8}
								color={useColorModeValue("gray.700", "whiteAlpha.900")}
								shadow="base"
							>
								<VStack spacing={5}>
									<FormControl isRequired>
										<FormLabel>Nome</FormLabel>

										<InputGroup>
											<InputLeftElement children={<BsPerson />} />
											<Input type="text" name="name" placeholder="Nome" />
										</InputGroup>
									</FormControl>

									<FormControl isRequired>
										<FormLabel>Email</FormLabel>

										<InputGroup>
											<InputLeftElement children={<MdOutlineEmail />} />
											<Input type="email" name="email" placeholder="Email" />
										</InputGroup>
									</FormControl>

									<FormControl isRequired>
										<FormLabel>Mensaxe</FormLabel>

										<Textarea
											name="message"
											placeholder="¡Escibe equí el to mensaxe!"
											rows={6}
											resize="none"
										/>
									</FormControl>

									<Button
										colorScheme="blue"
										bg="blue.400"
										color="white"
										_hover={{
											bg: "blue.500",
										}}
										isFullWidth
									>
										Envar Mensaxe
									</Button>
								</VStack>
							</Box>
						</Stack>
					</VStack>
				</Box>
			</Box>
		</Flex>
	);
};

export default Contact;
