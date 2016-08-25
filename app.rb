require 'sidekiq'

Sidekiq.configure_client do |config|
  config.redis = { :namespace => 'x', :size => 1, :url => 'redis://127.0.0.1:6379' }
end

class HardWorker
  include Sidekiq::Worker

  def perform(name="John", age=35)
  end
end

HardWorker.perform_async "John", 35