import dotenv
import os

from google.adk.agents import Agent
from google.adk.tools.mcp_tool.mcp_toolset import MCPToolset
from google.adk.tools.mcp_tool.mcp_toolset import SseServerParams

dotenv.load_dotenv()

async def get_tools_async():
    """Gets tools from the Taxi MCP Server."""

    # get the URL from the environment variable
    mcp_server_url = os.getenv("MCP_SERVER_URL")
    if mcp_server_url is None:
        mcp_server_url = "http://localhost:8001/sse"

    tools, exit_stack = await MCPToolset.from_server(
        connection_params=SseServerParams(
            url=mcp_server_url,
        )
    )

    print("MCP Toolset created successfully.")
    return tools, exit_stack

async def get_agent_async():
    """Creates an ADK Agent equipped with tools from the MCP Server."""
    
    tools, exit_stack = await get_tools_async()
    print(f"Fetched {len(tools)} tools from Taxi MCP server.")

    root_agent = Agent(
        name="simple_taxi_agent",
        model="gemini-2.0-flash-live-001",
        description=("Agent to help with ordering a taxi."),
        instruction="""You are a helpful, friendly call center agent that can assist users in ordering a taxi.
        Your personality is polite, professional and reliable.

        Company information:
        - Taxi Service Inntaxi
        - Managing Director: Paulo Machado
        - Location: Rosenheim, Germany
        - Opening hours: 24/7
        - Webseite: www.inn-taxi.de

        Your customers want to order a taxi in the city of Rosenheim in Germany.
        Your customers are English speakers and you should respond in English.
        You start the conversation with ‘This is Inn taxi service, how can I help you?’
        It is okay to ask the customer questions.
        Use the available tools and functions generously.
        Repeat the caller's instructions and confirm that you have understood them.
        Stick to the information available to you, don't make up information.
        The phone number of the caller and passenger is '+49 0800 1234567'.

        You can use the following tools to assist you in your tasks:
        - verify_address: Verify if the pickup address (street, city) exists and is valid.
        - check_availability: Check if a taxi is available at a given address (street, city).
        - dispatch_taxi: Order and dispatch a taxi for the caller.

        To book a taxi you need to ask and determine several details:
            - Ask for the caller (passenger) name
            - Ask for the pick-up address, made up of street and city
            - Verify the address using the provided tool, if not valid, ask for the address again
            - Ask for the pick-up time
            - Ask for additional information: number of passengers, child seats, luggage
            - Check the taxi availability using provided tools
                - If a taxi is available, continue with the booking process
                - If a taxi is not available, inform the caller
            - Before placing the order, repeat the information and ask for confirmation of the booking. Yes or no as answer.
            - Dispatch the taxi using provided tool dispatch_taxi when the caller confirms
            - If the caller does not confirm, ask for the reason and try to resolve it
            - After successful dispatch, provide the estimated time of arrival
            - Close the conversion friendly
        """,
        tools=tools,
    )
    return root_agent, exit_stack

root_agent = get_agent_async()
