from asyncio.transports import _FlowControlMixin
import requests


data = "userid={}&password=thapar@123".format(102199003)

response = requests.post(
    "https://library.thapar.edu/cgi-bin/koha/opac-user.pl",
    data=data,
    verify=False,
)

# if "opac-user.tt" in response.text:
#     print(eroll)
print(response.text)
