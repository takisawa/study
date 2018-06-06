require 'singleton'
require 'test/unit'

class MySingleton
  include Singleton
end

class TestMySingleton < Test::Unit::TestCase
  def test_new
    assert_raise( NoMethodError ) { MySingleton.new }
  end

  def test_instance
    assert_instance_of(MySingleton, MySingleton.instance)

    assert_same(MySingleton.instance, MySingleton.instance)
  end
end
