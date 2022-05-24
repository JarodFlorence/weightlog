<?php namespace Danijwilliams\Dinero\Requests;

use Danijwilliams\Dinero\Builders\Builder;
use Danijwilliams\Dinero\Utils\RequestBuilder;

class BookingRequestBuilder extends RequestBuilder
{
	public function __construct( Builder $builder )
	{
		$this->parameters['fields'] = 'Guid,Number,Timestamp';

		parent::__construct( $builder );
	}
}
