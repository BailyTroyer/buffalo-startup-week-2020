# Part 4 (DNS)

## Overview

DNS is one of those things that's not really technical but you just need to get done. You've probably heard of GoDaddy and Google Domains, but you can also easily purchase domains directly in AWS within Route53. I personally prefer this since you can easily interact with your domain, create hosted zones and TLS certs without having to import a domain from a 3rd party. 

For this example we're interacting with a domain `bflobox.com` which I personally bought on AWS for $12/year. 

## Buy a Domain (Route53)

Head over to **Route53 > Domains > Register Domain** search for an available domain, **Add to cart** and then press **continue**, enter your personal information (don't forget to enable privcy protection), press **continue** and then press **complete Order**. Welcome! You're now a proud owner of a domain!

**Note:** More detailed information on purchasing a domain on Route53 can be found [here](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-register.html).

## Create a Hosted Zone

Assuming you purchased a domain via Route53, you should have a hosted zone already created named after the domain you bought. For me its `bflobox.com`. This hosted zone is the catch all bucket for adding DNS entries to your domain name. You can add more hosted zones, add single entries and tons of more like performing DNS swaps between multiple hosts and much more!

In the next section we're going to create another hosted zone for our Kubernetes cluster, so keep this tab open.

## Create a TLS Cert

While we're doing DNS configuration we're probably going to want to create a certificate so we can run everything with `https` prepended. To do so we just need to press a few more buttons in the AWS console.

Navigate to the AWS service **Certificate Manager**, select **request certificate > Request a public certificate** and then enter the following 2 domain names:

* `*.domain.com` - for me it was `bflobox.com`
* `domain.com` - for me it was `*.bflobox.com`

The wildcard (*) simply allows us to attach a cert to any subdomains off-of `bflobox.com` if we want other domains such as `api.bflobox.com`, `www.bflobox.com`, `dev.bflobox.com`, etc.

Upon creating a cert, you're going to need to validate your certificate. Select the DNS option as opposed to an email address since it tends to be much quicker. 

## Create a DNS Entry

From there you're going to need to add an entry in Route53 for a random wacky validation entry used by Certificate Manager to validate the cert created. 

For me my entry looked like this `asdffdsfdafds.bflobox.com. ` with a value of `fdafdsa.fdsafd.acm-validations.aws.`

To add that entry go back to **Route53**, select **Hosted Zones > {YOUR HOSTED ZONE} > Create record**. Check **Simple Routing > Define simple record** and add your subdomain entry which was `asdffdsfdafds.bflobox.com.` or more specifically just `asdffdsfdafds` and then for **value/Route traffic to** choose IP address or another value depending on record type.

After that select **CNAME** and enter the value you were given back in Cert Manager which looked like `fdafdsa.fdsafd.acm-validations.aws.`

Select **Define Simple Record > Create Records** and then grab a coffee. This takes a few minutes. After that you should be all green and that cert is setup and ready-to-use. 
