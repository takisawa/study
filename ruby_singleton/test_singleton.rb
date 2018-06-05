require 'singleton'
require 'test/unit'

class MySingleton
  include Singleton
end

class TestMySingleton < Test::Unit::TestCase
  def test_new
    assert_raise( NoMethodError ) { MySingleton.new }
  end
end
