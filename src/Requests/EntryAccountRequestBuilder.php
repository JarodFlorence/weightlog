<?php

namespace Danijwilliams\Dinero\Requests;

use Danijwilliams\Dinero\Builders\Builder;
use Danijwilliams\Dinero\Utils\RequestBuilder;

class EntryAccountRequestBuilder extends RequestBuilder
{
	public function __construct( Builder $builder ) {
		$this->parameters['fields'] = 'Name,AccountNumber,VatCode,Category,IsHidden,IsDefaultSalesAccount';

		parent::__construct( $builder );
	}
}
