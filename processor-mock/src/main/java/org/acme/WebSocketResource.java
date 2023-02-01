package org.acme;

import io.quarkus.scheduler.Scheduled;

import org.eclipse.microprofile.config.inject.ConfigProperty;
import org.eclipse.microprofile.faulttolerance.Retry;
import org.eclipse.microprofile.reactive.messaging.Incoming;
import org.eclipse.microprofile.reactive.messaging.Message;
import org.jboss.logging.Logger;
import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;

import javax.enterprise.context.ApplicationScoped;
import javax.websocket.OnClose;
import javax.websocket.OnError;
import javax.websocket.OnOpen;
import javax.websocket.Session;
import javax.websocket.server.ServerEndpoint;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletionStage;
import java.util.concurrent.ConcurrentHashMap;

@ServerEndpoint("/api/metrics")
@ApplicationScoped
public class WebSocketResource {
    
    private static final Logger LOGGER = Logger.getLogger(WebSocketResource.class);

    @ConfigProperty(name = "version")
    private String processorVersion;

    private JSONParser parser = new JSONParser();
    private JSONObject json = new JSONObject();

    // @todo do we need to use a redis cluster instead of hashmap?
    Map<String, Session> activeSessions = new ConcurrentHashMap<>();
    List<String> sessionsToBeRemoved = new ArrayList<>();

    @Incoming("power_consumption")
    @Retry(delay = 10, maxRetries = 5)
    public CompletionStage<Void> consumePowerConsumption(Message<String> message) {
        broadcast(message.getPayload());
        return message.ack();
    }
    @Incoming("noise")
    @Retry(delay = 10, maxRetries = 5)
    public CompletionStage<Void> consumeNoise(Message<String> message) {
        broadcast(message.getPayload());
        return message.ack();
    }
    @Incoming("speed")
    @Retry(delay = 10, maxRetries = 5)
    public CompletionStage<Void> consumeSpeed(Message<String> message) {
        broadcast(message.getPayload());
        return message.ack();
    }
    @Incoming("temperature")
    @Retry(delay = 10, maxRetries = 5)
    public CompletionStage<Void> consumeTemperature(Message<String> message) {
        broadcast(message.getPayload());
        return message.ack();
    }

    @OnOpen
    public void onOpen(Session session) {
        activeSessions.put(session.getId(), session);
    }

    @OnClose
    public void onClose(Session session) {
        activeSessions.remove(session.getId());
    }

    @OnError
    public void onError(Session session, Throwable throwable) {
        activeSessions.remove(session.getId());
    }

    @Scheduled(every = "120s")
    public void cleanupInactiveWebSocketSessions() {
        sessionsToBeRemoved.forEach(key -> activeSessions.remove(key));
        sessionsToBeRemoved.clear();
    }

    private String parseMessage(String message) {

        try {
            json = (JSONObject) parser.parse(message);
            json.put("version", processorVersion);
        } catch (org.json.simple.parser.ParseException e) {
            LOGGER.error(e.getMessage());
        }
        
        return json.toString();
    }

    private void broadcast(String message) {
        String versionedMessage = parseMessage(message);
        activeSessions.forEach((key, value) -> value.getAsyncRemote().sendObject(versionedMessage, result -> {
            if (result.getException() != null) {
                sessionsToBeRemoved.add(key);
                LOGGER.debug("Unable to send message: " + result.getException());
            }
        }));
    }
}
