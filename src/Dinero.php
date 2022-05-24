<?php

namespace Danijwilliams\Dinero;

use GuzzleHttp\Exception\ClientException;
use GuzzleHttp\Exception\ServerException;
use Danijwilliams\Dinero\Builders\BookingBuilder;
use Danijwilliams\Dinero\Builders\ContactBuilder;
use Danijwilliams\Dinero\Builders\CreditnoteBuilder;
use Danijwilliams\Dinero\Builders\DepositAccountBuilder;
use Danijwilliams\Dinero\Builders\EntryAccountBuilder;
use Danijwilliams\Dinero\Builders\InvoiceBuilder;
use Danijwilliams\Dinero\Builders\PaymentBuilder;
use Danijwilliams\Dinero\Builders\ProductBuilder;
use Danijwilliams\Dinero\Builders\PurchaseVoucherBuilder;
use Danijwilliams\Dinero\Exceptions\DineroRequestException;
use Danijwilliams\Dinero\Exceptions\DineroServerException;
use Danijwilliams\Dinero\Requests\BookingRequestBuilder;
use Danijwilliams\Dinero\Requests\ContactRequestBuilder;
use Danijwilliams\Dinero\Requests\CreditnoteRequestBuilder;
use Danijwilliams\Dinero\Requests\DepositAccountRequestBuilder;
use Danijwilliams\Dinero\Requests\EntryAccountRequestBuilder;
use Danijwilliams\Dinero\Requests\InvoiceRequestBuilder;
use Danijwilliams\Dinero\Requests\PaymentRequestBuilder;
use Danijwilliams\Dinero\Requests\ProductRequestBuilder;
use Danijwilliams\Dinero\Requests\PurchaseVoucherRequestBuilder;
use Danijwilliams\Dinero\Utils\Request;

class Dinero
{
    protected $request;

    private $clientId;
    private $clientSecret;

    private $authToken;
    private $org;

    public function __construct($clientId, $clientSecret, $token = null, $org = null, $clientConfig = [])
    {
        $this->clientId = $clientId;
        $this->clientSecret = $clientSecret;
        $this->authToken = $token;
        $this->org = $org;

        $this->request = new Request($clientId, $clientSecret, $this->authToken, $this->org, $clientConfig);
    }

    public function setAuth($token, $org = null)
    {
        $this->authToken = $token;
        $this->org = $org;

        $this->request = new Request($this->clientId, $this->clientSecret, $this->authToken, $this->org);
    }

    public function getAuthToken()
    {
        return $this->authToken;
    }

    public function getAuthUrl()
    {
        return $this->request->getAuthUrl();
    }

    public function getOrgId()
    {
        return $this->org;
    }

    public function auth($apiKey, $orgId = null)
    {
        try {
            $response = json_decode($this->request->curl->post($this->request->getAuthUrl(), [
                'form_params' => [
                    'grant_type' => 'password',
                    'scope'      => 'read write',
                    'username'   => $apiKey,
                    'password'   => $apiKey,
                ],
            ])->getBody()->getContents());

            $this->setAuth($response->access_token, $orgId);

            return $response;
        } catch (ClientException $exception) {
            throw new DineroRequestException($exception);
        } catch (ServerException $exception) {
            throw new DineroServerException($exception);
        }
    }

    public function contacts()
    {
        return new ContactRequestBuilder(new ContactBuilder($this->request));
    }

    public function invoices()
    {
        return new InvoiceRequestBuilder(new InvoiceBuilder($this->request));
    }

	public function paymentsForInvoice($invoiceId)
	{
		return new PaymentRequestBuilder(new PaymentBuilder($this->request, "invoices/{$invoiceId}/payments"));
	}

	public function bookInvoice($invoiceId)
	{
		return new BookingRequestBuilder(new BookingBuilder($this->request, "invoices/{$invoiceId}/book"));
	}

    public function products()
    {
        return new ProductRequestBuilder(new ProductBuilder($this->request));
    }

    public function creditnotes()
    {
        return new CreditnoteRequestBuilder(new CreditnoteBuilder($this->request));
    }

    public function entryAccounts() {

        return new EntryAccountRequestBuilder(new EntryAccountBuilder($this->request));
    }

    public function depositAccounts() {

        return new DepositAccountRequestBuilder(new DepositAccountBuilder($this->request));
    }

    public function purchaseVouchers()
    {

        return new PurchaseVoucherRequestBuilder(new PurchaseVoucherBuilder($this->request));
    }
}
